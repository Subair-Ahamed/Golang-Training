package main

import (
    "fmt"
    "net/http"
    "myproject/cmd/webapp/routes"
)

func main() {
    router := routes.NewRouter()

    // Register a handler for static files (such as CSS, JS, and images)
    fs := http.FileServer(http.Dir("web/static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Register the router as the handler for all incoming HTTP requests
    http.Handle("/", router)

    // Start the HTTP server and listen for incoming requests
    fmt.Println("Server is running on port 8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Error:", err)
    }
}
