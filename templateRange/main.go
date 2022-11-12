package main

import (
	"net/http"
	"html/template"
)

type GroceryList []string

var tpl *template.Template
var g GroceryList

func main () {
	g = GroceryList{"milk","eggs","green beans", "cheese", "flour", "sugan", "brocoli"}
	tpl,_ = tpl.ParseGlob("template/*.html")
	http.HandleFunc("/list",listHandler)
	http.ListenAndServe(":8080",nil)
}


func listHandler(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "groceries.html",g)
}