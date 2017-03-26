package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gopkg.in/gomail.v2"

	_ "github.com/mattn/go-sqlite3"
)

type Member struct {
	Id       int    `gorm:"AUTO_INCREMENT" json:"id"`
	Username string `gorm:"not null;unique" json:"username"`
	TeamId   int    `gorm:"not null" json:"team_id"`
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
	db.Order("team_id").Find(&mems)

	c.JSON(200, mems)
}

func createMember(c *gin.Context) {
	db := initDB()
	defer db.Close()

	var mem Member
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
	// Email setup
	d := gomail.NewDialer("mail.securemail.vn", 25, os.Getenv("EMAIL"), os.Getenv("PASSWORD"))
	s, err := d.Dial()
	if err != nil {
		panic(err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL"))
	m.SetAddressHeader("To", "hoanghiepjp96@gmail.com", "Hiep")
	m.SetHeader("Subject", "Test")
	m.SetBody("text/plain", "Hello")

	db := initDB()
	defer db.Close()

	var chosen []Member
	db.Order("random()").Find(&chosen)

	// If number of player is odd, the last one won't play
	// 'coz the list is already randomized, it's totally fair!
	var dropMem Member
	if len(chosen)%2 != 0 {
		dropMem = chosen[len(chosen)-1]
		if dropMem.Id != 0 {
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

	if dropMem.Id != 0 {
		// Send email
		if err := gomail.Send(s, m); err != nil {
			panic(err)
		}

		// prepend dropMem to chosen
		c.JSON(200, append([]Member{dropMem}, chosen...))
	} else {
		c.JSON(200, chosen)
	}
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
		v1.PATCH("/draw", groupMembers)
	}

	router.Run(":8080")
}
