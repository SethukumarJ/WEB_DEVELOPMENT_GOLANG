package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template
var name = "John"

func main() {

	tpl, _ = tpl.ParseGlob("templates/*.html")
	http.HandleFunc("/welcome", welcomeHandler)
	http.ListenAndServe(":8080", nil)
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("indexHandler running")
	tpl.ExecuteTemplate(w, "welcome.html", name)
}
