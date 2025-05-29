package routes

import (
	"github.com/NicholasRaynes/tasktrack/backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/health", controllers.HealthCheck)
}
