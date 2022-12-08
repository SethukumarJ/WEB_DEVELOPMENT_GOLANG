package main

import (
	"log"
	"jwt/database"
	"jwt/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	router := gin.New()
	router.Use(gin.Logger())

	database.InitDB()

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.Run(":8080")
}