package helpers

import (
	"github.com/gin-gonic/gin"
)

func Error(err error) map[string]any {
	return gin.H{"error": err.Error()}
}
