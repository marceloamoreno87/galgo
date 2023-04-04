package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	Database "github.com/marceloamoreno87/galgo-api/config"
	Route "github.com/marceloamoreno87/galgo-api/routes"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Database.ConnectDatabase()
	router := gin.Default()
	Route.RegisterRoutes(router)
	router.Run(os.Getenv("PORT"))
}
