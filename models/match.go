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
	Team1Score int `sql:"DEFAULT:-1"`
	Team2Score int `sql:"DEFAULT:-1"`
}

type Score struct {
	Team1Score int `json:"score_team_1"`
	Team2Score int `json:"score_team_2"`
}

func (match *Match) GetMatchInfo(newMatch Match) {
	match.TournamentID = newMatch.TournamentID
	match.Team1ID = newMatch.Team1ID
	match.Team2ID = newMatch.Team2ID
	match.Team1Score = newMatch.Team1Score
	match.Team2Score = newMatch.Team2Score
}
