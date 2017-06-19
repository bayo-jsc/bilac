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
	Team1Score int `sql:"DEFAULT:0"`
	Team2Score int `sql:"DEFAULT:0"`
}

type Score struct {
	Team1Score int `json:"score_team_1"`
	Team2Score int `json:"score_team_2"`
}
