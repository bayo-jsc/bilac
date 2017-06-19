package main

import (
  "github.com/gin-gonic/gin"

  "./models"
)

func listMembers(c *gin.Context) {
  db := models.InitDB()
  defer db.Close()

  var mems []models.Member
  db.Find(&mems)

  c.JSON(200, mems)
}

func createMember(c *gin.Context) {
  db := models.InitDB()
  defer db.Close()

  var mem models.Member
  c.Bind(&mem)

  if mem.Username == "" {
    c.JSON(400, gin.H{"error": "Name not appropriate"})
  } else {
    if err := db.Create(&mem).Error; err != nil {
      c.JSON(500, gin.H{"error": err})
    } else {
      c.JSON(201, mem)
    }
  }
}

func showMember(c *gin.Context) {
  db := models.InitDB()
  defer db.Close()

  id := c.Params.ByName("id")
  var mem models.Member

  db.First(&mem, id)
  if mem.ID != 0 {
    c.JSON(200, mem)
  } else {
    c.JSON(404, gin.H{"error": "Member not found"})
  }
}

func updateMember(c *gin.Context) {
  db := models.InitDB()
  defer db.Close()

  id := c.Params.ByName("id")
  var mem models.Member

  db.First(&mem, id)
  if mem.ID == 0 {
    c.JSON(404, gin.H{"error": "Member not found"})
  } else {
    var uMem models.Member
    c.Bind(&uMem)

    if err := db.Model(&mem).Update("username", uMem.Username).Error; err != nil {
      c.JSON(400, gin.H{"error": err})
    } else {
      c.JSON(200, mem)
    }
  }
}

func destroyMember(c *gin.Context) {
  db := models.InitDB()
  defer db.Close()

  id := c.Params.ByName("id")
  var mem models.Member

  db.First(&mem, id)
  if mem.ID != 0 {
    if err := db.Delete(&mem).Error; err != nil {
      c.JSON(500, gin.H{"error": err})
    } else {
      c.Writer.WriteHeader(204)
    }
  } else {
    c.JSON(404, gin.H{"error": "Member not found"})
  }
}
