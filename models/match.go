package models

import (
	"github.com/jinzhu/gorm"
)

type Match struct {
	gorm.Model
	Tournament Tournament `gorm:"ForeignKey:TournamentID"`
	TournamentID uint
	Team1 Team `gorm:"ForeignKey:Team1ID"`
	Team1ID uint
	Team2 Team `gorm:"ForeignKey:Team2ID"`
	Team2ID uint
	Team1Score int
	Team2Score int
}
