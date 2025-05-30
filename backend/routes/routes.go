package routes

import (
	"github.com/NicholasRaynes/tasktrack/backend/controllers"
	"github.com/NicholasRaynes/tasktrack/backend/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/health", controllers.HealthCheck)
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/profile", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "You are authenticated 🎉"})
	})
}
