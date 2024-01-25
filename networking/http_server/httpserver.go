package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", greet) //syntax : http.HandleFunc(string pattern, func(wr snippet))

	http.HandleFunc("/intro", intro)

	http.ListenAndServe(":9000", nil) //output syntax: http.ListenAndServe(host address string, handler)
}

func greet(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1>Hi Prakash!, How are you?</h1>") //Fprintf(w,"write message to display")
}

func intro(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "This is Go HTTP Server !") //we can use http.ServeFile - (w ,r, "file to be displayed")

}
