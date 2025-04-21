// This defines the package name. "main" is special—it tells Go to build an executable program.
package main

// Import necessary packages
import (

	// Used for logging messages to the console (like printing errors or info).
	"log"

	// Provides HTTP client and server implementations for building web servers.
	"net/http"
)

func main() {

	http.HandleFunc("/receipts/process", ProcessReceiptHandler)
	http.HandleFunc("/receipts/", GetPointsHandler)

	// Print a message to the console saying the server is running on port 8080.
	log.Println("Server running on port 8080")

	// Start the HTTP server on port 8080.
	// If the server fails to start, log.Fatal will print the error and stop the program.
	log.Fatal(http.ListenAndServe(":8080", nil))
}
