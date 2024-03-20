package models

import "gorm.io/gorm"

type MessageForum struct {
	gorm.Model
	Content string `json:"content"`
	UserId  uint   `json:"userId"`
}
