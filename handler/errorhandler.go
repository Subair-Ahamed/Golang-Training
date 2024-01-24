package main

import (
	"fmt"
	"net/http"
)

func errorhandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Error message....!", http.StatusInternalServerError) //shows the internal server error
}
func main() {

	http.HandleFunc("/error", errorhandler)

	fmt.Println("Server running on port 7000")

	err := http.ListenAndServe(":7000", nil)

	if err != nil {
		fmt.Println("Error running server", err)
	}
}
