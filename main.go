package main

import (
	"net/http"
	"player-go-gin/logic"
	"player-go-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Poker Go Gin Player")
	})

	r.POST("/", func(c *gin.Context) {
		var table models.Table
		if c.Bind(&table) == nil {
			bet := logic.Decide(table)
			c.JSON(http.StatusOK, bet)
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "StatusInternalServerError"})
	})
	_ = r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
