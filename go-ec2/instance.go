package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		fmt.Println("Error creating session", err)
		return
	}

	// Create a new EC2 service client:
	svc := ec2.New(sess)

	// Set the key pair name:
	keyPairName := ""

	// Create the key pair:
	_, err = svc.CreateKeyPair(&ec2.CreateKeyPairInput{
		KeyName: aws.String(keyPairName),
	})
	if err != nil {
		fmt.Println("Error creating key pair", err)
		return
	}

	// Set the instance parameters:
	instanceParams := &ec2.RunInstancesInput{
		ImageId:      aws.String("ami-00c39f71452c08778"),
		InstanceType: aws.String("t2.micro"),
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1),
		KeyName:      aws.String(keyPairName), // Add the key pair name to the instance parameters
	}

	// Create the EC2 instance:
	resp, err := svc.RunInstances(instanceParams)
	if err != nil {
		fmt.Println("Error creating EC2 instance", err)
		return
	}

	// Get the instance ID:
	instanceID := resp.Instances[0].InstanceId
	fmt.Println("Created instance", *instanceID, "with key pair", keyPairName)

	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()

		// Get the instance ID, AMI ID, and instance type from the query params
		instanceID := queryParams.Get("instance_id")
		amiID := queryParams.Get("ami_id")
		instanceType := queryParams.Get("instance_type")

		// Create an AWS session
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("us-east-1"),
			// Credentials: credentials.NewStaticCredentials("your_access_key_id", "your_secret_access_key", ""),
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error creating AWS session: %v", err)
			return
		}

		// Create an EC2 client
		svc := ec2.New(sess)

		// Create input parameters for the DescribeInstances API call
		input := &ec2.DescribeInstancesInput{
			Filters: []*ec2.Filter{
				{
					Name: aws.String("instance-id"),
					Values: []*string{
						aws.String(instanceID),
					},
				},
				{
					Name: aws.String("image-id"),
					Values: []*string{
						aws.String(amiID),
					},
				},
				{
					Name: aws.String("instance-type"),
					Values: []*string{
						aws.String(instanceType),
					},
				},
			},
		}

		// Call the DescribeInstances API with the input parameters
		result, err := svc.DescribeInstances(input)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error calling DescribeInstances API: %v", err)
			return
		}

		// Build the response string
		var response strings.Builder
		for _, reservation := range result.Reservations {
			for _, instance := range reservation.Instances {
				response.WriteString(fmt.Sprintf("Instance ID: %s\n", *instance.InstanceId))
				response.WriteString(fmt.Sprintf("Instance Type: %s\n", *instance.InstanceType))
				response.WriteString(fmt.Sprintf("AMI ID: %s\n", *instance.ImageId))
				response.WriteString(fmt.Sprintf("Public IP Address: %s\n", *instance.PublicIpAddress))
				response.WriteString("\n")
			}
		}

		// Set the response headers and write the response string
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, response.String())
	})

	// Start the server and listen for requests
	fmt.Println("Server listening on port 9080")
	http.ListenAndServe(":9080", nil)
}
