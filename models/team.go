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
}
