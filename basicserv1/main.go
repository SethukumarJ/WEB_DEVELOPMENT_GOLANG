package main

import (

	"fmt" 
	"net/http"
)


func helloHandleFunc(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w,"Hello World !")
}

func main () {

	http.HandleFunc("/hello", helloHandleFunc)
	http.ListenAndServe(":8080",nil)

}