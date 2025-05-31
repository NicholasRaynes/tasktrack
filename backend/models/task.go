package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	ProjectID uint   `json:"project_id"`
}
