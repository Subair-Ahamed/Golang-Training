package main

import (
	"fmt"
	"net/http" 
)

func main(){

	//'http.HandleFunc' is a simpler way to register handlers using functions instead of objects.

	http.HandleFunc("/",greet) //syntax : http.HandleFunc(string address pattern, func_name(wr snippet))

	http.HandleFunc("/bye",sendoff)

	http.ListenAndServe(":9000",nil) //output syntax: http.ListenAndServe(host address string, handler)
}

func greet(w http.ResponseWriter, r *http.Request){ 
	fmt.Fprintf(w,"Hi Praksh!, How are you?") //Fprintf(w,"write message to display")
}

func sendoff(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Thank You for visting!...")
}