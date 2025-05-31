package controllers

import (
	"net/http"

	"github.com/NicholasRaynes/tasktrack/backend/database"
	"github.com/NicholasRaynes/tasktrack/backend/models"

	"github.com/gin-gonic/gin"
)

func CreateProject(c *gin.Context) {
	var project models.Project
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&project)
	c.JSON(http.StatusCreated, project)
}

func GetProjects(c *gin.Context) {
	var projects []models.Project
	database.DB.Preload("Tasks").Find(&projects)
	c.JSON(http.StatusOK, projects)
}
