package main

import (
	"fmt"
	"../models"
)

func main() {
	db := models.InitDB()
	defer db.Close()

	db.LogMode(false)

	// Reset elo of every member to 1000
	var members []models.Member
	db.Find(&members)

	fmt.Printf("Reset all members's elo to 1000\n")
	for _, mem := range members {
		fmt.Printf("Reseting member %s\n", mem.Username)
		mem.Elo = 1000
		db.Save(&mem)
	}

	var matches []models.Match
	db.Find(&matches)

	fmt.Printf("Update elo\n")
	for _, match := range matches  {
		fmt.Printf("Updating match %d\n", match.ID)
		match.UpdateElo(db)
	}
	fmt.Printf("Done!!!\n")
}
