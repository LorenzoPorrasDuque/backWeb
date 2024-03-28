package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type MessageForum struct {
	gorm.Model
	Content string `json:"content"`
	UserId  uint   `json:"userId"`
}

func GetAllYourPost(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, _ := c.Request.Cookie("Token")
		idUser := GetIdJWT(tokenString.Value)
		var messageForum []MessageForum
		db.Find(&messageForum, idUser)
		c.JSON(200, messageForum)
	}
}
func GetAllPost(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var messageForum []MessageForum
		db.Find(&messageForum)
		c.JSON(200, messageForum)
	}
}

func CreatePost(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var messageForum MessageForum
		c.BindJSON(&messageForum)

		//Esto se repite, deberia ser una funcion en el jwt, pero solo lo hace 3 veces asi que ......
		tokenString, _ := c.Request.Cookie("Token")
		idUser := GetIdJWT(tokenString.Value)
		id, _ := strconv.Atoi(idUser)
		messageForum.UserId = uint(id)

		db.Create(&messageForum)
		c.JSON(200, messageForum)
	}
}
func UpdatePost(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var messageForum MessageForum
		db.First(&messageForum, c.Param("id"))
		c.BindJSON(&messageForum)
		db.Save(&messageForum)
		c.JSON(200, messageForum)
	}
}
