package internal

import "github.com/gin-gonic/gin"

func HomeRoute(c *gin.Context) {
	c.HTML(200, "home.html", gin.H{
		"message": "Ping-Pong",
		"title":   "Home",
	})
}

func PongRoute(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
