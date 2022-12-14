package main

import (

	"net/http"
	"html/template"
)

var tpl *template.Template

func main () {

	//func ParseGlob(pattern string) (*Template, error)
	tpl = template.Must(template.ParseGlob("templates/*.html"))
	//func (t *Template) ParseFiles(filenames ...string) (*Template, error)
	// tpl, _ = tpl.ParseGlob("templates/*.html")


	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/login", loginHandler)
	http.ListenAndServe(":8080", nil)
	

}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w,"index.html", nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w,"about.html",nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w,"contact.html", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "login.html",nil)
}

