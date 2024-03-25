package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string `json:"username" gorm:"unique:not null:require"` //probar si esto funciona
	Password    string `json:"password"`
	Description string `json:"description"`
	UrlImage    string `json:"urlImage"`

	// Relationship
	MessagesForum []MessageForum `json:"messagesForum" gorm:"foreignKey:UserId"`
	Fighters      []Fighter      `json:"fighters" gorm:"many2many:UserId"`
	History       []History      `json:"history" gorm:"many2many:UserId"`
}

func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		c.BindJSON(&user)
		db.Create(&user)
		c.JSON(200, user)
	}
}
