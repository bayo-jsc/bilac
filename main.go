package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"./controllers"
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

	router.GET("/elo", func (c *gin.Context) {
		c.HTML(200, "elo.tpl", gin.H{})
	})

	// API v2 routers
	v2 := router.Group("/api/v2")
	{
		v2.GET("/members", controllers.ListMembers)
		v2.POST("/members", controllers.CreateMember)
		v2.GET("/members/:id", controllers.ShowMember)
		v2.PATCH("/members/:id", controllers.UpdateMember)
		v2.DELETE("/members/:id", controllers.DestroyMember)

		v2.GET("/tournaments", controllers.ListTournaments)
		v2.POST("/tournaments", controllers.CreateTournament)

		v2.GET("/tournaments/:id", controllers.GetTournament)
		v2.PATCH("/tournaments/:id/matches/:match_id", controllers.UpdateMatchScore)

		v2.PATCH("/tournaments/:id/shuffle", controllers.ShuffleMatch)

		//v2.GET("/members/:id/matches", getMemberMatches)
	}

	router.Run(":" + BILAC_PORT)
}
