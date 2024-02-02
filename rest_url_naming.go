package main

import (
	"fmt"
	"net/http"
)

func main() {

  //RESTful URI should refer to a resource that is a thing (noun) instead of referring to an action (verb) 
	//because nouns have properties that verbs do not have – similarly, resources have attributes.

	//Use lowercase letters for route names. Go's HTTP router is case-sensitive by default.
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/hello/", helloHandlerByID) // Add the new route here
	http.HandleFunc("/bye", byeHandler)
	http.HandleFunc("/bye/", byeHandlerByID) 
	http.HandleFunc("/bye/post", byeHandlerPOST) // Add the new POST route here

	//if you have nested resources, represent the hierarchy in the route. 
	//For example: /users/{userId}/comments 

	fmt.Println("Server is running on port 8000")
	http.ListenAndServe(":8000", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet { //this specifies the the method of api
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		return
	}

	w.WriteHeader(http.StatusOK) //if ok:
	fmt.Fprintf(w, "GET /hello - Hello, World!")
}

func byeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		return
	}
	w.WriteHeader(http.StatusOK) //if ok:
	fmt.Fprintf(w, "GET /bye - Goodbye, World!")
}

func helloHandlerByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		return
	}

	// Extract the ID from the URL path
  // if the request URL is "/bye/123", this line will extract "123" as the ID.
	helloID := r.URL.Path[len("/hello/"):]
	if helloID == "" {
		w.WriteHeader(http.StatusBadRequest) //It represents the status code 400, which indicates a bad request from the client.
		fmt.Fprintf(w, "Invalid request")
		return
	}

	w.WriteHeader(http.StatusOK) //if ok:
	fmt.Fprintf(w, "GET /hello/%s - Hello, %s!", helloID, helloID)
}

func byeHandlerByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		return
	}

	// Extract the ID from the URL path
	byeID := r.URL.Path[len("/bye/"):]
	if byeID == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request")
		return
	}

	w.WriteHeader(http.StatusOK) //if ok:
	fmt.Fprintf(w, "GET /bye/%s - Goodbye, %s!", byeID, byeID)
}

func byeHandlerPOST(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost { //here method will be POST
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		return
	}

	// Process the POST request here
	w.WriteHeader(http.StatusCreated) 
	fmt.Fprintf(w, "POST /bye - Create new bye entry")
}
