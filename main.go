package main

import (
	"backWeb/database"
	"backWeb/models"
	"fmt"
)

func init() {
	db := database.ConnectToDb()
	db.AutoMigrate(models.User{}, models.MessageForum{})

}

func main() {
	fmt.Print("Hello World")
}
