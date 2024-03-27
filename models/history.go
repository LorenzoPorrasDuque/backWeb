package models

import "gorm.io/gorm"

type History struct {
	gorm.Model

	Content string `json:"content"`
	Winner  int    `json:"winner"`
}
