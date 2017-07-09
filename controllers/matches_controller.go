package controllers

import (
	"github.com/gin-gonic/gin"
	"../models"
)

func ListMatches(c *gin.Context) {
	db := models.InitDB()
	defer db.Close()

	id := c.Params.ByName("id")
	var tour models.Tournament
	db.Preload("Matches").Find(&tour, id)

	c.JSON(200, tour.Matches)
}