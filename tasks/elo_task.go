package main

import (
	"fmt"
	"../models"
)

func main() {
	db := models.InitDB()
	defer db.Close()

	// Reset elo of every member to 1000
	var members []models.Member
	db.Find(&members)

	fmt.Printf("Reset all members's elo to 1000\n")
	for _, mem := range members {
		mem.Elo = 1000
		db.Save(&mem)
	}

	var matches []models.Match
	db.Find(&matches)

	fmt.Printf("Updating elo\n")
	for _, match := range matches  {
		match.UpdateElo(db)
	}
}
