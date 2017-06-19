package main

import (
	"os"
	"github.com/gin-gonic/gin"
)

var (
	BILAC_PORT string = os.Getenv("BILAC_PORT")
)

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

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "table.tpl", gin.H{})
	})

	router.GET("/draw", func (c *gin.Context) {
		c.HTML(200, "draw.tpl", gin.H{})
	})


	// API v1 routers
	v1 := router.Group("/api/v1")
	{
		v1.GET("/members", listMembers)
		v1.POST("/members", createMember)
		v1.GET("/members/:id", showMember)
		v1.PATCH("/members/:id", updateMember)
		v1.DELETE("/members/:id", destroyMember)

		v1.GET("/tournaments", listTournaments)
		v1.POST("/tournaments", createTournament)

		v1.GET("/last-tournament", lastTournament)
		v1.PATCH("/tournaments/:id/matches/:match_id", updateMatchScore)
	}

	router.Run(":" + BILAC_PORT)
}
