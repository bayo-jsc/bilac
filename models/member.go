package models

import (
	"github.com/jinzhu/gorm"
)

type Member struct {
	gorm.Model
	Username string `gorm:"not null;unique" json:"username"`
	Elo int `json:"elo" sql:"DEFAULT:1000"`
}

func (member Member) AddElo(amount int) {
	db := InitDB()
	defer db.Close()

	member.Elo += amount
	db.Save(&member)
}