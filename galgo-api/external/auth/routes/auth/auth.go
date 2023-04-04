package book

import (
	"github.com/gin-gonic/gin"
	AuthController "github.com/marceloamoreno87/galgo-api/external/auth/api/controllers"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/")
	routes.POST("/login", AuthController.Login)
	routes.POST("/logout", AuthController.Logout)
}
