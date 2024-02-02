package handlers

import (
	"fmt"
	"net/http"
)

// IndexHandler handles requests to the root endpoint "/"
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to My Go Project Structure!")
}
