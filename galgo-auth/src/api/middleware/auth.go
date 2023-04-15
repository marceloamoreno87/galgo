package middleware

import (
	"github.com/gin-gonic/gin"
)

func Run() gin.HandlerFunc {
	return func(c *gin.Context) {
		println("boa!")
		c.Next()
	}
}
