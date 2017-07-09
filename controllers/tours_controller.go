package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"../models"
)

func ListTournaments(c *gin.Context) {
	db := models.InitDB()
	defer db.Close()

	var tours []models.Tournament
	db.Order("created_at DESC").Find(&tours)

	c.JSON(200, tours)
}

func GetTournament(c *gin.Context) {
	db := models.InitDB()
	defer db.Close()

	id := c.Params.ByName("id")
	var tour models.Tournament

	db.Preload("Matches").Preload("Teams", func(db *gorm.DB) *gorm.DB {
		return db.Order("teams.points DESC, teams.gd DESC, teams.played_matches")
	}).Preload("Teams.Member1").Preload("Teams.Member2").Find(&tour, id)

	c.JSON(200, tour)
}

func CreateTournament(c *gin.Context) {
	db := models.InitDB()
	defer db.Close()

	var request models.TeamRequest

	c.BindJSON(&request)

	tour := models.Tournament{}
	db.Create(&tour)
	if err := db.Create(&tour).Error; err == nil {
		c.JSON(500, gin.H{"error": err})
	} else {
		// Create teams
		var teams []models.Team
		teams = tour.CreateTeams(db, request)

		// Create matches
		tour.CreateMatches(db, teams)

		c.JSON(201, tour)
	}
}

func UpdateMatchScore(c *gin.Context)  {
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
		team1.UpdateTeamScore(db)
		team2.UpdateTeamScore(db)
		match.UpdateElo(db)
		c.JSON(201, match)
	} else {
		c.JSON(500, gin.H{"error": err})
	}
}
