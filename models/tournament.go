package models

import (
	"github.com/jinzhu/gorm"
)

type Tournament struct {
	gorm.Model

	Matches []Match `gorm:"ForeignKey:TournamentID"`
	Teams []Team `gorm:"ForeignKey:TournamentID"`
}
