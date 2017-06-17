package models

import (
	"github.com/jinzhu/gorm"
)

type Match struct {
	gorm.Model
	Tournament Tournament `gorm:"ForeignKey:TournamentRefer"`
	Team1 Team `gorm:"ForeignKey:TeamRefer"`
	Team2 Team `gorm:"ForeignKey:TeamRefer"`
	Team1Score int
	Team2Score int
}
