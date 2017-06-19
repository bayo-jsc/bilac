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
	PlayedMatches int
	GF int
	GA int
	GD int
	Points int
}

type TeamRequest struct {
	Teams []struct {
		Member1_id int `json:"member1_id"`
		Member2_id int `json:"member2_id"`
	} `json:"teams"`
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
	team.GF, team.GA, team.Points, team.PlayedMatches = 0, 0, 0, 0

	var tour Tournament
	db.Model(team).Related(&tour)

	var matches []Match
	db.Model(tour).Where("team1_score >= ? OR team2_score >= ?", 0, 0).Related(&matches)

	for _, match := range matches  {
		if team.ID == match.Team1ID {
			team.GF += match.Team1Score
			team.GA += match.Team2Score
			team.Points += GetPoint(match.Team1Score, match.Team2Score)
			team.PlayedMatches += 1
		} else if team.ID == match.Team2ID {
			team.GF += match.Team2Score
			team.GA += match.Team1Score
			team.Points += GetPoint(match.Team2Score, match.Team1Score)
			team.PlayedMatches += 1
		}
	}
	team.GD = team.GF - team.GA
	db.Save(&team)
}
