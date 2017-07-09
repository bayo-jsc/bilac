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

	Mem1EloBefore int
	Mem1EloAfter int
	Mem2EloBefore int
	Mem2EloAfter int
	Mem3EloBefore int
	Mem3EloAfter int
	Mem4EloBefore int
	Mem4EloAfter int
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

func (match Match) UpdateElo(db *gorm.DB) {
	if match.Team1Score < 0 || match.Team2Score < 0 {
		return
	}

	var team1, team2 Team
	db.Model(match).Related(&team1, "Team1ID")
	db.Model(match).Related(&team2, "Team2ID")

	// Preload members
	team1.LoadMembers(db)
	team2.LoadMembers(db)

	// Assign before elo for members
	match.Mem1EloBefore = team1.Member1.Elo
	match.Mem2EloBefore = team1.Member2.Elo
	match.Mem3EloBefore = team2.Member1.Elo
	match.Mem4EloBefore = team2.Member2.Elo

	// Average elo of each teams
	elo1 := team1.AvgElo(db)
	elo2 := team2.AvgElo(db)

	// Expectation scores
	exp1 := 1 / float64(1 + math.Pow(10, (elo2 - elo1) / 500))
	exp2 := 1 / float64(1 + math.Pow(10, (elo1 - elo2) / 500))

	// Real scores
	var s1, s2 float64
	if match.Team1Score >= match.Team2Score {
		s1 = float64(match.Team1Score) / float64(match.Team1Score + match.Team2Score)
		s2 = 1 - s1
	} else {
		s1 = 1 - float64(match.Team2Score) / float64(match.Team1Score + match.Team2Score)
		s2 = 1 - s1
	}

	// Update team elo
	team1.UpdateElo(db, 100 * (s1 - exp1))
	team2.UpdateElo(db, 100 * (s2 - exp2))

	// Assign after elo
	team1.LoadMembers(db)
	team2.LoadMembers(db)
	match.Mem1EloAfter = team1.Member1.Elo
	match.Mem2EloAfter = team1.Member2.Elo
	match.Mem3EloAfter = team2.Member1.Elo
	match.Mem4EloAfter = team2.Member2.Elo
	db.Set("gorm:save_associations", false).Save(&match)
}
