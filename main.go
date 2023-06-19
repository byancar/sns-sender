package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sns-sender/app"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	router := gin.Default()

	router.POST("/message", func(c *gin.Context) {
		var message Message

		if err := c.ShouldBindJSON(&message); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		app.SendMessage(message.Text)
		c.JSON(http.StatusOK, gin.H{"message": message.Text})
	})

	router.Run(":8000")
}
