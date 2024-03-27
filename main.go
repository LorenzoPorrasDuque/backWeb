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
	r.GET("/login", models.Logger(db))
	r.POST("/createUser", models.CreateUser(db))

	//A partir de aqui todos los endpoints estan protegidos por token
	group := r.Group("/user", models.ValidateToken)

	group.GET("/getUser", models.GetUser(db))
	group.DELETE("/deleteUser", models.DeleteUser(db))
	group.PUT("/updateUser", models.UpdateUser(db))

	r.Run()

}
