package models

import (
	"github.com/jinzhu/gorm"
)

type Member struct {
	gorm.Model
	Username string `gorm:"not null;unique" json:"username"`
}
