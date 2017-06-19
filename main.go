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

	// API v2 routers
	v2 := router.Group("/api/v2")
	{
		v2.GET("/members", listMembers)
		v2.POST("/members", createMember)
		v2.GET("/members/:id", showMember)
		v2.PATCH("/members/:id", updateMember)
		v2.DELETE("/members/:id", destroyMember)

		v2.GET("/tournaments", listTournaments)
		v2.POST("/tournaments", createTournament)

		v2.GET("/last-tournament", lastTournament)
		v2.PATCH("/tournaments/:id/matches/:match_id", updateMatchScore)

		v2.PATCH("/tournaments/:id/shuffle", shuffleMatch)
	}

	router.Run(":" + BILAC_PORT)
}
