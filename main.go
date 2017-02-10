package main

import (
  "github.com/gin-gonic/gin"
  "github.com/jinzhu/gorm"

  _ "github.com/mattn/go-sqlite3"
)

type Member struct {
  Id int `gorm:"AUTO_INCREMENT" json:"id"`
  Username string `gorm:"not null;unique" json:"username"`
}

func initDB() *gorm.DB {
  db, err := gorm.Open("sqlite3", "db/table_football.db")
  if err != nil {
    panic(err)
  }

  db.LogMode(true)
  if !db.HasTable(&Member{}) {
    db.CreateTable(&Member{})
    db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Member{})
  }
  return db
}

func listMembers(c *gin.Context) {
  db := initDB()
  defer db.Close()

  var mems []Member
  db.Find(&mems)

  c.JSON(200, mems)
}

func createMember(c *gin.Context) {
  db := initDB()
  defer db.Close()

  var mem Member
  c.Bind(&mem)

  if err := db.Create(&mem).Error; err != nil {
    c.JSON(500, gin.H{"error": "Something's wrong"})
  } else {
    c.JSON(201, mem)
  }
}

func showMember(c *gin.Context) {
  db := initDB()
  defer db.Close()

  id := c.Params.ByName("id")
  var mem Member

  db.First(&mem, id)
  if mem.Id != 0 {
    c.JSON(200, mem)
  } else {
    c.JSON(404, gin.H{"error": "Member not found"})
  }
}

func main() {
  router := gin.Default()

  // Routers config
  router.GET("/members", listMembers)
  router.POST("/members", createMember)
  router.GET("/members/:id", showMember)

  router.Run(":8080")
}
