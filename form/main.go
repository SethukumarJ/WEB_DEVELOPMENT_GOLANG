package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Sub struct {
	Username string
	Data     string
}

var tpl *template.Template

func main() {

	tpl, _ = tpl.ParseGlob("templates/*.html")
	http.HandleFunc("/getform", getFormHandler)
	http.HandleFunc("/processget", processgetHandler)
	http.HandleFunc("/postform", postFormHandler)
	http.HandleFunc("/processpost", processPostHandler)
	http.ListenAndServe(":8080", nil)
}

func getFormHandler(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "getform.html", nil)
}

func processgetHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("ProcessGetHandler running")

	var s Sub

	s.Username = r.FormValue("username")
	s.Data = r.FormValue("data")
	tpl.ExecuteTemplate(w, "thanks.html", s)
}

func processPostHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("ProcessGetHandler running")

	var s Sub

	s.Username = r.FormValue("username")
	s.Data = r.FormValue("data")
	tpl.ExecuteTemplate(w, "thanks.html", s)
}

func postFormHandler(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "postform.html", nil)
}
