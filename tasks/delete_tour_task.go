package main

import (
	"os"
	"fmt"
	"../models"
)
func main()  {
	db := models.InitDB()
	defer db.Close()

	id := os.Args[1]

	fmt.Printf("Deleting tournament with ID = %s\n", id)
	var tour models.Tournament
	db.Find(&tour, id)

	if !db.NewRecord(tour) {
		tour.Delete(db)
	}
	fmt.Printf("Done!!!")
}
