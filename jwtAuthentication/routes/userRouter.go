package routes

import (
	controller "jwt/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.LoadHTMLGlob("templates/*.html")
	incomingRoutes.GET("/wadmin", controller.Wadmin)
	incomingRoutes.GET("/home", controller.Home)
	incomingRoutes.GET("/delete/:name", controller.DeleteUser)
	incomingRoutes.POST("/update/:name", controller.UpdateUser)
	incomingRoutes.POST("/create", controller.CreateUser)
}