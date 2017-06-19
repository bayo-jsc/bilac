package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"./models"
)

func listTournaments(c *gin.Context) {
	db := models.InitDB()
	defer db.Close()

	var tours []models.Tournament
	db.Order("CreatedAt").Find(&tours)

	c.JSON(200, tours)
}

func lastTournament(c *gin.Context) {
	db := models.InitDB()
	defer db.Close()

	var tour models.Tournament
	db.Order("created_at desc").Preload("Matches").Preload("Teams", func(db *gorm.DB) *gorm.DB {
		return db.Order("teams.points DESC, teams.GD DESC")
	}).Preload("Teams.Member1").Preload("Teams.Member2").First(&tour)

	c.JSON(200, tour)
}

func createTournament(c *gin.Context) {
	db := models.InitDB()
	defer db.Close()

	var request struct {
		Teams []struct {
			Member1_id int `json:"member1_id"`
			Member2_id int `json:"member2_id"`
		} `json:"teams"`
	}

	c.BindJSON(&request)
	teams_raw := request.Teams

	tour := models.Tournament{}
	db.Create(&tour)
	if err := db.Create(&tour).Error; err == nil {
		c.JSON(500, gin.H{"error": err})
	} else {
		// Create teams
		for i := 0; i < len(teams_raw); i++ {
			var member1 models.Member
			var member2 models.Member

			db.Find(&member1, teams_raw[i].Member1_id)
			db.Find(&member2, teams_raw[i].Member2_id)

			team := models.Team {
				Tournament: tour,
				Member1: member1,
				Member2: member2,
			}
			db.Create(&team)
		}

		// Create matches
		var teams []models.Team
		db.Model(&tour).Related(&teams)
		for i := 0; i < len(teams)-1; i++ {
			for j := i + 1; j < len(teams); j++ {
				match := models.Match {
					Tournament: tour,
					Team1: teams[i],
					Team2: teams[j],
				}

				db.Create(&match)
			}
		}
		c.JSON(201, tour)
	}
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
		c.JSON(201, match)
	} else {
		c.JSON(500, gin.H{"error": err})
	}
}