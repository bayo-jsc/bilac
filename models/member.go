package models

import (
	"github.com/jinzhu/gorm"
)

type Member struct {
	gorm.Model
	Username string `gorm:"not null;unique" json:"username"`
	Elo int `json:"elo" sql:"DEFAULT:1000"`
}

func (member Member) AddElo(db *gorm.DB, amount int) {
	member.Elo += amount
	db.Save(&member)
}