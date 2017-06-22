package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"bilac/models"
)

func listTournaments(c *gin.Context) {
	db := models.InitDB()
	defer db.Close()

	var tours []models.Tournament
	db.Order("created_at DESC").Find(&tours)

	c.JSON(200, tours)
}

func getTournament(c *gin.Context) {
	db := models.InitDB()
	defer db.Close()

	id := c.Params.ByName("id")
	var tour models.Tournament

	db.Preload("Matches").Preload("Teams", func(db *gorm.DB) *gorm.DB {
		return db.Order("teams.points DESC, teams.gd DESC, teams.played_matches")
	}).Preload("Teams.Member1").Preload("Teams.Member2").Find(&tour, id)

	c.JSON(200, tour)
}

func createTournament(c *gin.Context) {
	db := models.InitDB()
	defer db.Close()

	var request models.TeamRequest

	c.BindJSON(&request)
	teams_raw := request.Teams

	tour := models.Tournament{}
	db.Create(&tour)
	if err := db.Create(&tour).Error; err == nil {
		c.JSON(500, gin.H{"error": err})
	} else {
		// Create teams
		for i := 0; i < len(teams_raw); i++ {
			team := models.Team {
				TournamentID: tour.ID,
				Member1ID: uint(teams_raw[i].Member1_id),
				Member2ID: uint(teams_raw[i].Member2_id),
			}
			db.Create(&team)
		}

		// Create matches
		var teams []models.Team
		db.Model(&tour).Related(&teams)
		for i := 0; i < len(teams)-1; i++ {
			for j := i + 1; j < len(teams); j++ {
				match := models.Match {
					TournamentID: tour.ID,
					Team1ID: teams[i].ID,
					Team2ID: teams[j].ID,
				}

				db.Create(&match)
			}
		}
		// Shuffle Match
		db.Model(tour).Related(&tour.Matches)
		tour.ShuffleMatch()
		c.JSON(201, tour)
	}
}

func shuffleMatch(c *gin.Context) {
	db := models.InitDB()
	defer db.Close()

	tour_id := c.Params.ByName("id")
	var tour models.Tournament
	db.Find(&tour, tour_id)

	tour.ShuffleMatch()

	c.JSON(200, tour)
}

func updateMatchScore(c *gin.Context)  {
	db := models.InitDB()
	defer db.Close()

	tour_id := c.Params.ByName("id")
	var tour models.Tournament
	db.Find(&tour, tour_id)

	match_id := c.Params.ByName("match_id")
	var match models.Match
	db.Preload("Team1").Preload("Team2").Find(&match, match_id)

	var score models.Score
	c.BindJSON(&score)

	match.Team1Score = score.Team1Score
	match.Team2Score = score.Team2Score

	if err := db.Save(&match).Error; err == nil {
		team1 := match.Team1
		team2 := match.Team2
		team1.UpdateTeamScore()
		team2.UpdateTeamScore()
		match.UpdateElo()
		c.JSON(201, match)
	} else {
		c.JSON(500, gin.H{"error": err})
	}
}
