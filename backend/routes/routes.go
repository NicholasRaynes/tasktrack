package routes

import (
	"tasktrack/backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/health", controllers.HealthCheck)
}
