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

	//User endpoints

	group.GET("/getUser", models.GetUser(db))
	group.GET("/getOtherUser/:id", models.GetOtherUser(db))
	group.GET("/getAllUser", models.GetAllUser(db))
	group.DELETE("/deleteUser", models.DeleteUser(db))
	group.PUT("/updateUser", models.UpdateUser(db))

	//Figther endpoints

	group.GET("/getFighter/:id", models.GetFighter(db))
	group.GET("/getYourFighters/", models.GetYourFighters(db))
	group.GET("/getAllFighters", models.GetAllFighters(db))
	group.POST("/createFighter", models.CreateFighter(db))
	group.DELETE("/deleteFighter/:id", models.DeleteFighter(db)) //Problema, este parcero borra cualquiera no solo los de la persona
	group.PUT("/updateFighter/:id", models.UpdateFighter(db))    //Mismo problema anterior

	//MessageForum endpoints
	group.GET("/getAllYourPost", models.GetAllYourPost(db))
	group.GET("/getAllPost", models.GetAllPost(db))
	group.POST("/createPost", models.CreatePost(db))
	group.PUT("/updatePost/:id", models.UpdatePost(db))

	//History endpoints
	group.POST("/createHistory/:id1/:id2", models.CreateHistory(db))
	group.GET("/getAllHistories", models.SearchAllHistories(db))
	group.GET("/getHistory/:id", models.SearchHistory(db))
	r.Run()

}
