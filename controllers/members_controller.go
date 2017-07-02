package controllers

import (
	"github.com/gin-gonic/gin"

	"../models"
)

func ListMembers(c *gin.Context) {
	db := models.InitDB()
	defer db.Close()

	sort := c.Query("sort")

	var mems []models.Member
	db.Order(sort).Find(&mems)

	c.JSON(200, mems)
}

func CreateMember(c *gin.Context) {
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

func ShowMember(c *gin.Context) {
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

func UpdateMember(c *gin.Context) {
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

func DestroyMember(c *gin.Context) {
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

//func getMemberMatches(c *gin.Context) {
//	db := models.InitDB()
//	defer db.Close()
//
//	id := c.Params.ByName("id")
//	var mem models.Member
//	db.First(&mem, id)
//
//	var teamIDs []int
//	db.Table("teams").Where("member1_id = ? OR member2_id = ?", id, id).Pluck("ID", &teamIDs)
//
//	var matches []models.Match
//	//var matches []struct {
//	//	ID int
//	//
//	//}
//	db.Where("team1_id in (?)", teamIDs).Or("team2_id in (?)", teamIDs).Find(&matches)
//	for _, match := range matches {
//
//	}
//	c.JSON(200, matches)
//}
