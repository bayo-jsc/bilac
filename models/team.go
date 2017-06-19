package models

import (
	"github.com/jinzhu/gorm"
)

type Team struct {
	gorm.Model
	Tournament Tournament `gorm:"ForeignKey:TournamentID"`
	TournamentID uint
	Member1 Member `gorm:"ForeignKey:Member1ID"`
	Member1ID uint
	Member2 Member `gorm:"ForeignKey:Member2ID"`
	Member2ID uint
	GF int
	GA int
	GD int
	Points int
}

func GetPoint(x, y int) int {
	if x > y {
		return 3
	}
	if x < y {
		return 0
	}
	return 1
}

func (team Team) UpdateTeamScore() {
	db := InitDB()
	team.GF, team.GA, team.Points = 0, 0, 0

	var tour Tournament
	db.Model(team).Related(&tour)

	var matches []Match
	db.Model(tour).Where("updated_at > created_at").Related(&matches)

	for _, match := range matches  {

			if team.ID == match.Team1ID {
				team.GF += match.Team1Score
				team.GA += match.Team2Score
				team.Points += GetPoint(match.Team1Score, match.Team2Score)
			} else if team.ID == match.Team2ID {
				team.GF += match.Team2Score
				team.GA += match.Team1Score
				team.Points += GetPoint(match.Team2Score, match.Team1Score)
			}
	}
	team.GD = team.GF - team.GA
	db.Save(&team)
}
