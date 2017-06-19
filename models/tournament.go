package models

import (
	"github.com/jinzhu/gorm"
)

type Tournament struct {
	gorm.Model

	Matches []Match `gorm:"ForeignKey:TournamentID"`
	Teams []Team `gorm:"ForeignKey:TournamentID"`
}

func (tour *Tournament) ShuffleMatch() {
  db := InitDB()
  defer db.Close()

  var oldOrder []Match
  db.Model(tour).Related(&oldOrder)

  db.Model(tour).Order("RANDOM()").Related(&tour.Matches)

  for i, _ := range tour.Matches {
    tour.Matches[i].GetMatchInfo(oldOrder[i])
  }

  db.Save(&tour)
  return
}
