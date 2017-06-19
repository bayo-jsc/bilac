package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"

	"../models"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "db/table_football.db")
	if err != nil {
		panic(err)
	}

	db.LogMode(true)
	// Create table
	if !db.HasTable(&models.Member{}) {
		db.CreateTable(&models.Member{})
	}

	if !db.HasTable(&models.Tournament{}) {
		db.CreateTable(&models.Tournament{})
	}

	if !db.HasTable(&models.Match{}) {
		db.CreateTable(&models.Match{})
	}

	if !db.HasTable(&models.Team{}) {
		db.CreateTable(&models.Team{})
	}

	return db
}
