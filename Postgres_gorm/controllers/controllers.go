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
