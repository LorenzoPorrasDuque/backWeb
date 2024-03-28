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

func CreateHistory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var history History
		c.BindJSON(&history)
		var fighter1, fighter2 Fighter
		db.First(&fighter1, c.Param("id1"))
		db.First(&fighter2, c.Param("id2"))
		history.Fighter = append(history.Fighter, &fighter1)
		history.Fighter = append(history.Fighter, &fighter2)
		db.Create(&history)
	}
}
