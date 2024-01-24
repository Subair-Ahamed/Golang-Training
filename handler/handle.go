package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello this is handlers in go!...")
}

func main() {

	handler := &MyHandler{}
	http.Handle("/", handler)
	
	fmt.Println("Server running on port 7000")

	err := http.ListenAndServe(":7000", nil)

	if err != nil {
		fmt.Println("Error running server", err)
	}
}
