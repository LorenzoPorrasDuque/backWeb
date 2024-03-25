package models

import "gorm.io/gorm"

type History struct {
	gorm.Model
	UserId   []User    `json:"userId" gorm:"many2many:history_user;"`
	Fighters []Fighter `json:"fighterId" gorm:"many2many:history_fighter;"`
	Content  string    `json:"content"`
	Winner   int       `json:"winner"`
}
