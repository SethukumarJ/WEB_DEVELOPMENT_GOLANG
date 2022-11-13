package main

import (

	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"unicode"

	"golang.org/x/crypto/bcypt"
	_ "github.com/go-sql-driven/mysql"

)


var tpl *template.Template
var db *sql.DB

func main (){

	tpl= template.Must(template.ParseGlob("templates/*.html"))

	var err error

	db, err  = sql.Open("mysql", "root:password@tcp(localhost:3306)/testdb")
	if err != nil {
		
	}

}