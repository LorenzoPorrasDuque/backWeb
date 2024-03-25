package main

import (
	"backWeb/database"
	"backWeb/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	db := database.ConnectToDb()
	db.AutoMigrate(&models.User{}, &models.MessageForum{}, &models.Fighter{}, &models.History{})

}

func main() {
	fmt.Print("Hello World")
	r := gin.Default()

	group := r.Group("/api", Logger)

	group.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()

}
