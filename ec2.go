package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func main() {
	// Load the shared AWS configuration file
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		fmt.Println("Error loading AWS configuration:", err)
		os.Exit(1)
	}

	// Create a new EC2 service client
	svc := ec2.New(cfg)

	// Set up the parameters for the instance
	input := &ec2.RunInstancesInput{
		ImageId:      aws.String(""), // Amazon Linux 2 AMI (HVM), SSD Volume Type
		InstanceType: ec2.InstanceTypeT2Micro,
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1),
	}

	// Launch the instance
	result, err := svc.RunInstances(context.Background(), input)
	if err != nil {
		fmt.Println("Error launching EC2 instance:", err)
		os.Exit(1)
	}

	// Print the instance ID
	fmt.Println("Instance ID:", *result.Instances[0].InstanceId)
}
