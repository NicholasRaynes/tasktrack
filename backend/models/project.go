package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Title  string `json:"title"`
	UserID uint   `json:"user_id"`
	Tasks  []Task `json:"tasks"`
}
