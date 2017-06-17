package main

import (
	"os"
	"github.com/gin-gonic/gin"
	//"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"

	"./models"
	"github.com/jinzhu/gorm"
	//"fmt"
)

var (
	BILAC_PORT string = os.Getenv("BILAC_PORT")
)

func initDB() *gorm.DB {
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

func listMembers(c *gin.Context) {
	db := initDB()
	defer db.Close()

	var mems []models.Member
	db.Find(&mems)

	c.JSON(200, mems)
}

func createMember(c *gin.Context) {
	db := initDB()
	defer db.Close()

	var mem models.Member
	c.Bind(&mem)

	if mem.Username == "" {
		c.JSON(400, gin.H{"error": "Name not appropriate"})
	} else {
		if err := db.Create(&mem).Error; err != nil {
			c.JSON(500, gin.H{"error": "Something's wrong"})
		} else {
			c.JSON(201, mem)
		}
	}
}

func showMember(c *gin.Context) {
	db := initDB()
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
	db := initDB()
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
	var mem models.Member

	db.First(&mem, id)
	if mem.ID != 0 {
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

	var chosen []models.Member
	db.Order("random()").Find(&chosen)

	// If number of player is odd, the last one won't play
	// 'coz the list is already randomized, it's totally fair!
	var dropMem models.Member
	if len(chosen)%2 != 0 {
		dropMem = chosen[len(chosen)-1]
		if dropMem.ID != 0 {
			db.Model(&dropMem).Update("team_id", nil)
		}
		chosen = chosen[:len(chosen)-1]
	}

	// init group id with 0
	// for each 2-player (start with 0) increase group id by 1
	g := 0
	for k, _ := range chosen {
		if k%2 == 0 {
			g += 1
		}
		db.Model(&chosen[k]).Update("team_id", g)
	}

	if dropMem.ID != 0 {
		// prepend dropMem to chosen
		c.JSON(200, append([]models.Member{dropMem}, chosen...))
	} else {
		c.JSON(200, chosen)
	}
}

func serveFE(c *gin.Context) {
	c.HTML(200, "index.tpl", gin.H{})
}

func listTournaments(c *gin.Context) {
	db := initDB()
	defer db.Close()

	var tours []models.Tournament
	db.Order("CreatedAt").Find(&tours)

	c.JSON(200, tours)
}

func createTournament(c *gin.Context) {
	db := initDB()
	defer db.Close()

	var request struct {
		Teams []struct {
			Member1_id int `json:"member1_id"`
			Member2_id int `json:"member2_id"`
		} `json:"teams"`
	}

	c.BindJSON(&request)
	teams := request.Teams

	tour := models.Tournament{}
	db.Create(&tour)
	if err := db.Create(&tour).Error; err == nil {
		c.JSON(500, gin.H{"error": "Something's wrong"})
	} else {
		for i := 0; i < len(teams); i++ {
			var member1 models.Member
			var member2 models.Member

			db.Find(&member1, teams[i].Member1_id)
			db.Find(&member2, teams[i].Member2_id)

			team := models.Team {
				Tournament: tour,
				Member1: member1,
				Member2: member2,
			}
			db.Create(&team)
			c.JSON(201, team)
			return
		}
		c.JSON(201, tour)
	}
}

func showTeam(c *gin.Context) {
	db := initDB()
	defer db.Close()

	id := c.Params.ByName("id")
	var team models.Team
	db.Find(&team, id)

	c.JSON(200, team)
}

func init() {
	if BILAC_PORT == "" {
		BILAC_PORT = "8080"
	}
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
		v1.PATCH("/draw", groupMembers)

		v1.GET("/tournaments", listTournaments)
		v1.POST("/tournaments", createTournament)

		v1.GET("/teams/:id", showTeam)
	}

	router.Run(":" + BILAC_PORT)
}
