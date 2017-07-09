package models

import (
	"github.com/jinzhu/gorm"
	//"fmt"
	"math"
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

func getPoint(x, y int) int {
	if x > y {
		return 3
	}
	if x < y {
		return 0
	}
	return 1
}

func (team Team) UpdateTeamScore(db *gorm.DB) {
	team.GF, team.GA, team.Points, team.PlayedMatches = 0, 0, 0, 0

	var tour Tournament
	db.Model(team).Related(&tour)

	var matches []Match
	db.Model(tour).Where("team1_score >= ? OR team2_score >= ?", 0, 0).Related(&matches)

	for _, match := range matches  {
		if team.ID == match.Team1ID {
			team.GF += match.Team1Score
			team.GA += match.Team2Score
			team.Points += getPoint(match.Team1Score, match.Team2Score)
			team.PlayedMatches += 1
		} else if team.ID == match.Team2ID {
			team.GF += match.Team2Score
			team.GA += match.Team1Score
			team.Points += getPoint(match.Team2Score, match.Team1Score)
			team.PlayedMatches += 1
		}
	}
	team.GD = team.GF - team.GA
	db.Save(&team)
}

func (team *Team) LoadMembers(db *gorm.DB) {
	db.Preload("Member1").Preload("Member2").Find(&team, team.ID)
}

func (team Team) AvgElo(db *gorm.DB) float64 {
	if &team.Member1 == nil || &team.Member2 == nil {
		db.Preload("Member1").Preload("Member2").Find(&team, team.ID)
	}

	return float64(team.Member1.Elo + team.Member2.Elo) / 2
}

func (team Team) UpdateElo(db *gorm.DB, elo float64) {
	db.Preload("Member1").Preload("Member2").Find(&team, team.ID)
	smallRatio := math.Min(float64(team.Member1.Elo), float64(team.Member2.Elo)) /
						float64(team.Member1.Elo + team.Member2.Elo)

	var team1Ratio float64
	if elo >= 0 {
		if team.Member1.Elo >= team.Member2.Elo {
			team1Ratio = smallRatio
		} else {
			team1Ratio = 1 - smallRatio
		}
	} else {
		if team.Member1.Elo >= team.Member2.Elo {
			team1Ratio = 1 - smallRatio
		} else {
			team1Ratio = smallRatio
		}
	}

	//team1Ratio := 0.5
	//fmt.Printf("%.2f %.2f\n", elo, team1Ratio)
	team.Member1.AddElo(int(2 * elo * team1Ratio))
	team.Member2.AddElo(int(2 * elo * (1 - team1Ratio)))
}
