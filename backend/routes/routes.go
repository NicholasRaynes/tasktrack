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

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())

	// Project routes
	protected.POST("/projects", controllers.CreateProject)
	protected.GET("/projects", controllers.GetProjects)

	// Task routes
	protected.POST("/tasks", controllers.CreateTask)
	protected.GET("/tasks", controllers.GetTasks)
}
