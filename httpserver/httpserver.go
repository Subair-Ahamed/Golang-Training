package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", greet) //syntax : http.HandleFunc(string pattern, func(wr snippet))

	http.HandleFunc("/text", settype)

	http.HandleFunc("/bye", sendoff)

	http.ListenAndServe(":9000", nil) //output syntax: http.ListenAndServe(host address string, handler)
}

func greet(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1>Hi Prakash!, How are you?</h1>") //Fprintf(w,"write message to display")
}

func settype(w http.ResponseWriter, r *http.Request) { //w is the response and r is the request

	w.Header().Set("Content-Type", "text/html") //we can set the type of response displayed(html,text,etc...) usimg w.Header().Set(key string,value string)
	fmt.Fprintf(w, "<h1>These are http server's features</h1>")
	fmt.Fprintf(w, "<p>There are various methods and types which are defined in the package build a HTTP service. </p>")
}

func sendoff(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "index.html") //ServeFile - to display a file
}
