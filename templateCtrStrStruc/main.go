package main

import (
	"net/http"
	"html/template"
)

type productSpec struct {

	Size 	string 
	Weight 	float32
	Descr 	string

}

type product struct {

	ProductID	int
	Name		string
	Cost		float32
	Specs		productSpec

}

var tpl *template.Template
var product1 product

func main () {

	product1 = product {
		ProductID: 15,
		Name: "wicked cool Phone",
		Cost:899,
		Specs: productSpec {
			Size: "259 x 70 x 7 mm",
			Weight : 65,
			Descr: "Over priced shiny trhing desighned to shatte aon impoct",
		},
	}

	tpl,  _ = tpl.ParseGlob("templates/*.html")
	http.HandleFunc("/productinfo", productInfoHandler)
	http.ListenAndServe(":8080",nil)

}

func productInfoHandler(w http.ResponseWriter, r *http.Request) {

	tpl.ExecuteTemplate(w, "productinfo.html", product1)
}