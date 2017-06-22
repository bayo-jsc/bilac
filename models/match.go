package models

import (
	"github.com/jinzhu/gorm"
	"math"
	//"fmt"
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

func (match Match) UpdateElo() {
	db := InitDB()
	defer db.Close()

	var team1, team2 Team
	db.Model(match).Related(&team1, "Team1ID")
	db.Model(match).Related(&team2, "Team2ID")

	elo1 := team1.AvgElo()
	elo2 := team2.AvgElo()

	exp1 := 1 / float64(1 + math.Pow(10, (elo2 - elo1) / 500))
	exp2 := 1 / float64(1 + math.Pow(10, (elo1 - elo2) / 500))

	var s1, s2 float64
	if match.Team1Score >= match.Team2Score {
		s1 = float64(match.Team1Score) / float64(match.Team1Score + match.Team2Score)
		s2 = 1 - s1
	} else {
		s1 = 1 - float64(match.Team2Score) / float64(match.Team1Score + match.Team2Score)
		s2 = 1 - s1
	}

	//fmt.Printf("%.3f : %.3f : %.3f : %.3f\n", exp1, s1, exp2, s2)
	team1.UpdateElo(100 * (s1 - exp1))
	team2.UpdateElo(100 * (s2 - exp2))
}