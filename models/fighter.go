package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type Fighter struct {
	gorm.Model
	Name   string `json:"name" gorm:"not null"`
	Age    int    `json:"age" gorm:"not null"`
	UserId uint   `json:"userId" gorm:"not null"`
	Stats  Stats  `json:"stats" gorm:"embedded;not null"`
	// Relationship
	History []*History `json:"history" gorm:"many2many:fighter_history;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Stats struct {
	Strength int     `json:"strength"`
	Height   int     `json:"height"`
	Weight   int     `json:"weight"`
	Agility  int     `json:"agility"`
	Luck     float32 `json:"luck"`
}

func GetFighter(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var fighter Fighter
		db.Preload("History").First(&fighter, c.Param("id"))
		c.JSON(200, fighter)
	}
}
func GetYourFighters(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, _ := c.Request.Cookie("Token")
		idUser := GetIdJWT(tokenString.Value)
		var fighter []Fighter
		db.Find(&fighter, "user_id = ?", idUser)
		c.JSON(200, fighter)
	}
}
func GetAllFighters(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var fighter []Fighter
		db.Find(&fighter)
		c.JSON(200, fighter)
	}
}
func CreateFighter(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var fighter Fighter
		c.BindJSON(&fighter)
		tokenString, _ := c.Request.Cookie("Token")
		idUser := GetIdJWT(tokenString.Value)
		id, _ := strconv.Atoi(idUser)
		fighter.UserId = uint(id)
		db.Create(&fighter)
		c.JSON(200, fighter)
	}
}
func DeleteFighter(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var fighter Fighter
		db.First(&fighter, c.Param("id"))
		db.Delete(&fighter)
		c.JSON(200, fighter)
	}
}
func UpdateFighter(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		db.First(&user, c.Param("id"))
		c.BindJSON(&user)
		db.Save(&user)
		c.JSON(200, user)
	}
}
