package main

import (
	"fmt"
	"net/http"
)


// handler function that responds to client http requests
func helloHandleFunc(w http.ResponseWriter , r *http.Request) {

	// formats using the default formats for its operands and writes to w
	fmt.Fprint(w, "Hello World!")
}

func aboutHandleFunc(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "r.URL.path:, %s!", r.URL.Path)
}


func main() {
	// registers the handler funtion for the given pattern in the DefaultServeMux
	http.HandleFunc("/hello", helloHandleFunc)
	http.HandleFunc("/about", aboutHandleFunc)
	//listenAndServe listens on hte TCP network address addr and then calls Serve
	// with handler to handle requiests on incoming connections
	http.ListenAndServe("localhost:8080", nil)
	// ServeMux is an HTTP request multiplexer , It matches the URL of each incoming 
	// requet againset a list of registered patterns and calls the handler for the 
	// pattern theat most closely matches the URL.
}