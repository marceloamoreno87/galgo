package routes

import (
	"github.com/gin-gonic/gin"
	Book "github.com/marceloamoreno87/galgo-api/routes/books"
	User "github.com/marceloamoreno87/galgo-api/routes/users"
)

func RegisterRoutes(router *gin.Engine) {

	User.Routes(router)
	Book.Routes(router)
}
