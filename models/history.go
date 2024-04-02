package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type History struct {
	gorm.Model

	Content string `json:"content"`
	Winner  int    `json:"winner"`

	Fighter []*Fighter `json:"history" gorm:"many2many:fighter_history;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

//FALTA UPDATE DE HISTORIA VER EL MUCHOS A MUCHOS

func calculateWinner(f1, f2 Fighter) int {
	var winner int
	promFighter1 := (f1.Stats.Agility + f1.Stats.Strength + f1.Stats.Weight) / 3
	promFighter2 := (f2.Stats.Agility + f2.Stats.Strength + f2.Stats.Weight) / 3
	if promFighter1 > promFighter2 {
		winner = int(f1.ID)
	} else {
		winner = int(f2.ID)

	}
	return winner
}

func CreateHistory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		//Funciona, pero no tengo ningun sistema de verificacion para las varibales y que no sean nulas
		var history History
		c.BindJSON(&history)
		var fighter1, fighter2 Fighter
		db.First(&fighter1, c.Param("id1"))
		db.First(&fighter2, c.Param("id2"))
		history.Fighter = append(history.Fighter, &fighter1)
		history.Fighter = append(history.Fighter, &fighter2)
		history.Winner = calculateWinner(fighter1, fighter2)
		db.Create(&history)
	}
}
func SearchAllHistories(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var history []History
		db.Find(&history)
		c.JSON(200, history)
	}
}
func SearchHistory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var history History
		db.First(&history, c.Param("id"))
		c.JSON(200, history)
	}
}
