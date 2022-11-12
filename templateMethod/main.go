package main

import (

	"fmt"
	"html/template"
	"net/http"

)


var tpl *template.Template
type Price float64

func (p Price) CanCashPr() string {

	remainder := int(p*100) % 5
	quotient := int(p*100) / 5

	if remainder < 3 {
		pr := float64(quotient*5) / 100
		s := fmt.Sprintf("%.2f", pr)
		return s
	}

	pr := (float64(quotient*5) + 5) / 100
	s := fmt.Sprintf("%.2f", pr)
	return s
}


var p Price 

func main () {

	p = 3.94
	tpl,_ = tpl.ParseFiles("index.html")
	http.HandleFunc("/", indexHandler)
}

func indexHandler (w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", p)
}