package routes

import (
    "go-project/api/handlers"
    "net/http"
)

// StartServer initializes the HTTP server and starts listening
func StartServer() {
    http.HandleFunc("/", handlers.IndexHandler)
    http.ListenAndServe(":8080", nil)
}
