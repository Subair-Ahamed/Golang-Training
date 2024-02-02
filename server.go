package main

import (
	"fmt"
	"net/http"
)

// APIBaseURL represents the base URL of the API.
//const - names should be capitalized or fully uppercase
const API_URL = "https://api.example.com"

//struct - names should be with an uppercase letter, and capitalize the first letter of each subsequent word in the name.
type Handler struct{}

// ServeHTTP handles HTTP requests.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the API server at %s", API_URL)
}

func main() {
	// Create a new instance of Handler
	handler := &Handler{}

	// Register the handler to handle requests for the root route
	http.Handle("/net", handler)

	// Start the HTTP server on port 8080
	fmt.Printf("Server listening on %s:8080",API_URL)
	http.ListenAndServe(":8080", nil)
}
