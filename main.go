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

func updateMember(c *gin.Context) {
  db := initDB()
  defer db.Close()

  id := c.Params.ByName("id")
  var mem Member

  db.First(&mem, id)
  if mem.Id == 0 {
    c.JSON(404, gin.H{"error": "Member not found"})
  } else {
    var uMem Member
    c.Bind(&uMem)

    if err := db.Model(&mem).Update("username", uMem.Username).Error; err != nil {
      c.JSON(500, gin.H{"error": "Something's wrong"})
    } else {
      c.JSON(200, mem)
    }
  }
}

func destroyMember(c *gin.Context) {
  db := initDB()
  defer db.Close()

  id := c.Params.ByName("id")
  var mem Member

  db.First(&mem, id)
  if mem.Id != 0 {
    if err := db.Delete(&mem).Error; err != nil {
      c.JSON(500, gin.H{"error": "Something's wrong"})
    } else {
      c.Writer.WriteHeader(204)
    }
  } else {
    c.JSON(404, gin.H{"error": "Member not found"})
  }
}

func groupMembers(c *gin.Context) {
  db := initDB()
  defer db.Close()

  var chosen []Member
  db.Order("random()").Find(&chosen)

  // If number of player is odd, remove last one
  // 'coz the list is already randomized
  if len(chosen) % 2 != 0 {
    chosen = chosen[:len(chosen)-1]
  }

  c.JSON(200, chosen)
}

func serveFE(c *gin.Context) {
  c.HTML(200, "index.tpl", gin.H{})
}


func main() {
  router := gin.Default()

  // Normal routers
  router.LoadHTMLGlob("templates/*.tpl")
  router.Static("node_modules", "./node_modules")
  router.Static("static", "./static")
  router.GET("/", serveFE)

  // API v1 routers
  v1 := router.Group("/api/v1")
  {
    v1.GET("/members", listMembers)
    v1.POST("/members", createMember)
    v1.GET("/members/:id", showMember)
    v1.PATCH("/members/:id", updateMember)
    v1.DELETE("/members/:id", destroyMember)
    v1.GET("/draw", groupMembers)
  }

  router.Run(":8080")
}
