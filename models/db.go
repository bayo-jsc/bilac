package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	
)

func InitDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "db/table_football.db")
	if err != nil {
		panic(err)
	}

	db.LogMode(true)
	// Create table
	if !db.HasTable(&Member{}) {
		db.CreateTable(&Member{})
	}

	if !db.HasTable(&Tournament{}) {
		db.CreateTable(&Tournament{})
	}

	if !db.HasTable(&Match{}) {
		db.CreateTable(&Match{})
	}

	if !db.HasTable(&Team{}) {
		db.CreateTable(&Team{})
	}

	return db
}
