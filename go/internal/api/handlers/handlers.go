package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestFunc() {
	log.Println("Hello from api handlers")
}

func MovieHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}
