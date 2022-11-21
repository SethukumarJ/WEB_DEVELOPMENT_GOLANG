package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Person struct {
	gorm.Model

	Name  string
	Email string
	Books []Book
}

type Book struct {
	gorm.Model

	Title      string
	Auther     string
	CallNumber int
	PersonID   int
}

var db *gorm.DB
var err error

func main() {

	// Loading environmant variables

	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	// Database connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s dbPort=%s", host, user, dbName, password, dbPort)

	//openning connection to database

	db, err = gorm.Open(dialect, dbURI)
	if err != nil {
		log.Fatal(err)
	} else {

		fmt.Println("successfully connected!")

	}

}
