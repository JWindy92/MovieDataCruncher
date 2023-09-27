package server

import (
	"log"

	"github.com/JWindy92/MovieDataCruncher/config"
	"github.com/JWindy92/MovieDataCruncher/internal/api/handlers"
	"github.com/gin-gonic/gin"
)

func Run() {
	log.Println("Starting API Server")
	r := gin.Default()
	r.GET("/movie", handlers.MovieHandler())

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func TestFunc() {
	log.Println("Hello from api server")

	log.Println(config.Config.Messaging.Host)
}
