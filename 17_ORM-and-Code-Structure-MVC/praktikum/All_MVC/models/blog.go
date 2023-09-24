package models

import (
	"github.com/jinzhu/gorm"
)

type Blog struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	User    User
}