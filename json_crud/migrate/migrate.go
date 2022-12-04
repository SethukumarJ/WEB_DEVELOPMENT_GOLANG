package main

import (
	"jsc/initializers"
	"jsc/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()

}

func main() {

	initializers.DB.AutoMigrate(&models.Post{})
}
