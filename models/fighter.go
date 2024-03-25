package models

import "gorm.io/gorm"

type Fighter struct {
	gorm.Model
	Name   string `json:"name"`
	Age    int    `json:"age"`
	IsDead bool   `json:"isDead"`
	UserId uint   `json:"userId"`
	Stats  Stats  `json:"stats" gorm:"embedded"`
	// Relationship
	History []History `json:"history" gorm:"many2many:history_fighter;"`
}

type Stats struct {
	Strength int     `json:"strength"`
	Height   int     `json:"height"`
	Weight   int     `json:"weight"`
	Agility  int     `json:"agility"`
	Luck     float32 `json:"luck"`
}
