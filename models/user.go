package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string `json:"username" gorm:"unique;not null"` //probar si esto funciona
	Password    string `json:"password" gorm:"not null"`
	Description string `json:"description"`
	UrlImage    string `json:"urlImage"`

	// Relationship
	MessagesForum []*MessageForum `json:"messagesForum" gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL; " `
	Fighters      []*Fighter      `json:"fighters" gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		c.BindJSON(&user)
		db.Create(&user)
		c.JSON(200, user)
	}
}
func GetUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, _ := c.Request.Cookie("Token")
		idUser := GetIdJWT(tokenString.Value)
		var user User
		db.Preload("MessageForum").Preload("Fighters").Preload("Fighters.History").First(&user, idUser)
		c.JSON(200, user)
	}
}
func GetAllUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user []User
		db.Preload("MessageForum").Preload("Fighters").Find(&user)
		c.JSON(200, user)
	}
}
func GetOtherUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		db.Preload("MessageForum").Preload("Fighters").First(&user, c.Param("id"))
		c.JSON(200, user)
	}
}
func DeleteUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, _ := c.Request.Cookie("Token")
		idUser := GetIdJWT(tokenString.Value)
		var user User
		db.First(&user, idUser)
		db.Delete(&user)
		c.JSON(200, user)
	}
}
func UpdateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, _ := c.Request.Cookie("Token")
		idUser := GetIdJWT(tokenString.Value)
		var user User
		db.First(&user, idUser)
		c.BindJSON(&user)
		db.Save(&user)
		c.JSON(200, user)
	}
}
