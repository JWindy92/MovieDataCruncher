package server

import (
	"log"

	"github.com/JWindy92/MovieDataCruncher/config"
)

func TestFunc() {
	log.Println("Hello from api server")

	log.Println(config.Config.Messaging.Host)
	// messaging.PublishMessage()
	// messaging.ConsumeMessage()
}
