package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {

	//func parseFiles(filenames ...string) (*Template, error)
	// tpl = template.Must(template.ParseFiles("index1.html"))
	tpl = template.Must(template.ParseFiles("data1/index2.html"))
	// tpl, _ = template.ParseFiles("data1/data2/index3.html")
	// tpl, _ = template.ParseFiles("../index4.html")
	// func (t *Template) ParseFiles(filenames ...string) (*Template, error)
	//tpl, _ = tpl.ParseFiles("../index4.html")

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// func (t *Template) Execute(wr io.Writer, data interface{}) error
	tpl.Execute(w, nil)

}
