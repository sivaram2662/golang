package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		// Define a simple struct that holds some information
		type Info struct {
			Name  string `json:"yourname"`
			SName string `json:"sername"`
		}

		// Create a new instance of the Info struct
		info := Info{
			Name:  "siva",
			SName: "Duggireddy",
		}

		// Convert the Info struct to JSON
		jsonData, err := json.Marshal(info)
		if err != nil {
			http.Error(w, "Error converting data to JSON", http.StatusInternalServerError)
			return
		}

		// Set the response headers to indicate that the response contains JSON data
		w.Header().Set("Content-Type", "application/json")

		// Write the JSON data to the response
		w.Write(jsonData)
	})

	// Start the web server
	http.ListenAndServe(":6080", nil)
}
