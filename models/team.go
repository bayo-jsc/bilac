package models

import (
  "github.com/jinzhu/gorm"
)

type Team struct {
  gorm.Model
  Tournament Tournament `gorm:"ForeignKey:TournamentRefer"`
  Member1 Member `gorm:"ForeignKey:MemberRefer"`
  Member2 Member `gorm:"ForeignKey:MemberRefer"`
}
