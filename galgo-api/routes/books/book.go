package book

import (
	"github.com/gin-gonic/gin"
	BookController "github.com/marceloamoreno87/galgo-api/api/books/controllers"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/books")
	routes.GET("", BookController.Index)
	routes.POST("", BookController.Store)
	routes.GET("/:id", BookController.Show)
	routes.PUT("/:id", BookController.Update)
	routes.DELETE("/:id", BookController.Destroy)
}
