package main

import (
	"os"

	"github.com/NicholasRaynes/tasktrack/backend/config"
	"github.com/NicholasRaynes/tasktrack/backend/database"
	"github.com/NicholasRaynes/tasktrack/backend/routes"

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
