package main

import (
	"backWeb/database"
	"backWeb/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.ConnectToDb()
	db.AutoMigrate(&models.User{}, &models.MessageForum{}, &models.Fighter{}, &models.History{})
	fmt.Println("Hello World")

	r := gin.Default()
	// Enpoints para crear usuario y logearse
	r.GET("/login", Logger(db))
	r.POST("/createUser", models.CreateUser(db))

	//A partir de aqui todos los endpoints estan protegidos por token
	group := r.Group("/user", ValidateToken)

	group.GET("/get", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})

	})

	r.Run()

}
