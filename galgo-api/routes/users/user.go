package user

import (
	"github.com/gin-gonic/gin"
	UserController "github.com/marceloamoreno87/galgo-api/api/users/controllers"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/users")
	routes.GET("", UserController.Index)
	routes.POST("", UserController.Store)
	routes.GET("/:id", UserController.Show)
	routes.PUT("/:id", UserController.Update)
	routes.DELETE("/:id", UserController.Destroy)
}
