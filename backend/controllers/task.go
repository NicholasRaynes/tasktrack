package controllers

import (
	"net/http"

	"github.com/NicholasRaynes/tasktrack/backend/database"
	"github.com/NicholasRaynes/tasktrack/backend/models"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&task)
	c.JSON(http.StatusCreated, task)
}

func GetTasks(c *gin.Context) {
	var tasks []models.Task
	database.DB.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}
