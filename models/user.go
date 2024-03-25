package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Description string `json:"description"`
	UrlImage    string `json:"urlImage"`

	// Relationship
	MessagesForum []MessageForum `json:"messagesForum" gorm:"foreignKey:UserId"`
	Fighters      []Fighter      `json:"fighters" gorm:"many2many:UserId"`
	History       []History      `json:"history" gorm:"many2many:UserId"`
}
