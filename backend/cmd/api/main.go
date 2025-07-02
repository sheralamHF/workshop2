package main

import (
	"log"
	"net/http"
)

// AI generated
func main() {
	// Define the status endpoint handler
	http.HandleFunc("/status", statusHandler)
	
	// Start the server on port 8080
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// AI generated
func statusHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type and status code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	// Write the response
	w.Write([]byte(`{"status": true}`))
	
	log.Println("Status endpoint called")
} 