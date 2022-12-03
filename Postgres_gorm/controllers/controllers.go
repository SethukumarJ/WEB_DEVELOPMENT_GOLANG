package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"project_login/database"
	"project_login/helpers"
	"project_login/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var Store = sessions.NewCookieStore([]byte("secret"))

type Users struct {
	ID       int
	Username string
	Password string
}
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}


func VerifyPassword(userPassword string, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	if err != nil {
		check = false
	}
	return check

}