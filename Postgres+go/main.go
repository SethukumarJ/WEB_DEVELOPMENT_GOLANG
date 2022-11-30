package main 


import (

	"fmt"  //This is provided by go
	"database/sql"
	"log"
	_ "github.com/lib/pq" //Third party package

)

// Provide credentials for our database.

const (

	host = "localhost"
	port = 5432
	user = "assets"
	password = "1111"
	dname = "postgres_demo"

)


func main () {
	dsn := fmt.Sprintf("host =%s port =%d user = %s, password=%s, dname=%s",host,port,user,password,dname)

	//Establesh a connection ato the database
	db, err := sql.Open("postgres",dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()//Always do this before exiting
	//Test our connetion
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	
	//Let's insert a quote

	insertQuote := `
	INSERT INTO quotations (author_name, category, quote)
	VALUES ($1, $2, $3)
	RETURNING quotation_id
	`

	   quotation_id := 0
	err = db.QueryRow(insertQuote, "Lao Tzu",
 			"Life",
			"Mastering others is strenth. Mastering yourself is true power.").Scan(&quotation_id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The recently inserted record has quotation_id:", quotation_id)

}