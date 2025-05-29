package main

import (
	"os"
	"tasktrack/backend/config"
	"tasktrack/backend/database"
	"tasktrack/backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.Connect()

	r := gin.Default()
	routes.RegisterRoutes(r)

	port := os.Getenv("PORT")
	r.Run(":" + port)
}
