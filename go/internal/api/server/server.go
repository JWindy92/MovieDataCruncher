package server

import (
	"log"
	"net/http"

	"github.com/JWindy92/MovieDataCruncher/config"
	"github.com/gin-gonic/gin"
)

func Run() {
	log.Println("Starting API Server")
	r := gin.Default()
	r.GET("/movie", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
func TestFunc() {
	log.Println("Hello from api server")

	log.Println(config.Config.Messaging.Host)
	// messaging.PublishMessage()
	// messaging.ConsumeMessage()
}
