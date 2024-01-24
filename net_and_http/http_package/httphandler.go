package main

import (
    "fmt"
    "net/http"
)

type MyHandler struct{}

// ServeHTTP is the method required to implement the http.Handler interface

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {  //method syntax : func(receiver type) methodname(arguments){}

    fmt.Fprintf(w, "Hello, this is a simple HTTP handler!")
}

func main() {
    // Create an instance of MyHandler
    handler := &MyHandler{}

    //http.Handle registers a handler for a given route pattern. It takes in a route pattern and an object that implements the http.Handler interface
	
    http.Handle("/hello", handler)

    // Start the HTTP server on port 8080
    fmt.Println("Server listening on port 8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}
