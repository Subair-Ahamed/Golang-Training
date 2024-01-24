package main

import (
	"fmt"
	"net/http"
)

func display1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>HandleFunc() function</h1>")
}
func display2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello this is HandldeFunc in go handlers....!")
}
func main() {

	http.HandleFunc("/show1", display1)

	http.HandleFunc("/show2", display2)

	fmt.Println("Server running on port 7000")

	err := http.ListenAndServe(":7000", nil)

	if err != nil {
		fmt.Println("Error running server", err)
	}
}
