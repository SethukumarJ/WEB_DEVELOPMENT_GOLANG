package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"unicode"

	
	"golang.org/x/crypto/bcrypt"
	"github.com/go-sql-driver/mysql"
	
	) 
// "github.com/go-sql-driver/mysql v1.6.0" is not a valid import path
var tpl *template.Template
var db *sql.DB

func main() {

	tpl = template.Must(template.ParseGlob("templates/*.html"))

	var err error

	db, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/testdb")
	if err != nil {

		panic(err.Error())

	}
	defer db.Close()
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/registerauth", registerAuthHandler)
	http.ListenAndServe("localhost:8080", nil)

}

// registerhandler serves form for registering new users
func registerHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("**********RegisterHandler running*********")
	tpl.ExecuteTemplate(w, "register.html", nil)
}

func registerAuthHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("*************registrationHandler running******")
	r.ParseForm()
	username := r.FormValue("username")

	nameAlphaNumeric := true

	for _, char := range username {

		if unicode.IsLetter(char) == false && unicode.IsNumber(char) == false {
			nameAlphaNumeric = false
		}
		
	}
	//check username pswdlength
	var nameLength bool

	if 5 <= len(username) && len(username) <= 50 {
		nameLength = true
	}

	password := r.FormValue("password")
	fmt.Println("password:", password, "\npswdLength:", len(password))
	//variables that mush pass for password creation criteria
	var pswdLowercase, pswdUppercase, pswdNumber,pswdSpecial,pswdLength,paswdNoSpaces bool
	pswdNoSpaces = true

	for _, char := range password {

		switch {
		case unicode.IsLower(char):
			pswdLowercase = true
			
		case unicode.IsUpper(char):
			pswdUppercase = true

		case unicode.IsNumber(char):
			pswdNumber = true

		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			pswdSpecial = true

		case unicode.IsSpace(int32(char)):
			paswdNoSpaces = false
		}
	}


	if 11 < len(password) && len(password) < 60 {
		pswdLength = true
	}

	fmt.Println("pswdLowercase : ",pswdLowercase, "\npswdUppercase :", pswdUppercase, "\npswdNumber :", pswdNumber,
"\npswdSpecial :", pswdSpecial, "\npasswordNoSpaces :",paswdNoSpaces)

if !pswdLowercase || !pswdUppercase || !pswdNumber || !pswdLength || !pswdLength || !paswdNoSpaces {
	tpl.ExecuteTemplate(w , "register.html","please check username and password criteria")
	return
}

//check if username already exists for availability

stmt := "SELECT UserId From bcrypt Where username = ?"

row := db.QueryRow(stmt, username)

var uID string
err := row.Scan(&uID)
if err != sql.ErrNoRows {
	fmt.Println("username already exists, err: ", err)
	tpl.ExecuteTemplate(w, "register.html","username already taken")
	return
}


var hash []byte

hash, err = bcrypt.GenerateFromPassWord([]byte(password),bcrypt.DefaultCost)
if err != nil {
	fmt.Println("bcrypt err: ",err)
}
}
