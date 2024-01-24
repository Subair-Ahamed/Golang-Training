//middleware is typically implemented as functions or handlers that wrap the main HTTP handler.

package main

import (
	"fmt"
	"net/http"
	"time"
)

//http.HandlerFunc is a type used to wrap a function, converting it into an http.Handler type.
//It is typically used when you need to use a function as a handler with functions like http.Handle.

func middleware(middle http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { 

		startTime := time.Now()

        //middle.ServeHTTP(w, r) line is calling the ServeHTTP method of the next handler in the chain. 
		//This is crucial for the middleware to allow the request to proceed to the main handler
		middle.ServeHTTP(w, r)
        
		//the middleware logs information about the request, such as the HTTP method, request URI, and the time
		fmt.Printf("[%v] %v %v\n", r.Method, r.RequestURI, time.Since(startTime))
	})
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Midlleware Handlers...!")
}

func main() {

	http.Handle("/mid", middleware(http.HandlerFunc(hello)))

	fmt.Println("Server listening on port 7000")

	err := http.ListenAndServe(":7000", nil)

	if err != nil {
		fmt.Println("Error occured", err)
	}
}
