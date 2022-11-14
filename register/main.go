package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"unicode"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
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

	var nameAlphaNumeric = true

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
	var pswdLowercase, pswdUppercase, pswdNumber, pswdSpecial, pswdLength, paswdNoSpaces bool
	paswdNoSpaces = true

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
	//Check pswd criteria

	fmt.Println("pswdLowercase : ", pswdLowercase, "\npswdUppercase :", pswdUppercase, "\npswdNumber :", pswdNumber,
		"\npswdSpecial :", pswdSpecial, "\npasswordNoSpaces :", paswdNoSpaces)

	if !pswdLowercase || !pswdUppercase || !pswdNumber || !pswdLength || !paswdNoSpaces || !nameAlphaNumeric || !nameLength {
		tpl.ExecuteTemplate(w, "register.html", "please check username and password criteria")
		return
	}

	//check if username already exists for availability

	stmt := "SELECT UserId From bcrypt Where username = ?"

	row := db.QueryRow(stmt, username)

	var uID string
	err := row.Scan(&uID)
	if err != sql.ErrNoRows {
		fmt.Println("username already exists, err: ", err)
		tpl.ExecuteTemplate(w, "register.html", "username already taken")
		return
	}

	var hash []byte

	hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("bcrypt err: ", err)
		tpl.ExecuteTemplate(w, "register.html", "there was a problem registering account")
		return
	}

	fmt.Println("hash:", hash)
	fmt.Println("string(hash):", string(hash))

	//func (db *DB) Prepare(qurery string) (*Stmt, error)

	var insertStmt *sql.Stmt

	insertStmt, err = db.Prepare("INSERT INTO bcrypt (Username, Hash) VALUES (?, ?);")
	if err != nil {
		fmt.Println("error preparing statement:", err)
		tpl.ExecuteTemplate(w, "regiseter.html", "there was a problem registering account")
		return
	}

	defer insertStmt.Close()
	var result sql.Result
	// func (s *Stmt) Exec(args ...interface{}) (Result, error)
	result, err = insertStmt.Exec(username, hash)
	rowsAff, _ := result.RowsAffected()
	lastIns, _ := result.LastInsertId()
	fmt.Println("rowAff:", rowsAff)
	fmt.Println("lastIns:", lastIns)
	fmt.Println("err:", err)
	if err != nil {
		fmt.Println("error inserting new user")
		tpl.ExecuteTemplate(w, "refister.html", "there was a problem registering account")
		return
	}
	fmt.Fprint(w, "congrats, yoiur account has been successfully created")
}
