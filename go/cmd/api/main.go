package main

import (
	"log"

	"github.com/JWindy92/MovieDataCruncher/config"
	"github.com/JWindy92/MovieDataCruncher/internal/api/server"
	"github.com/JWindy92/MovieDataCruncher/messaging"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		// Handle error, e.g., log and exit
		panic(err)
	}
	server.Run()
}

func testMsg() {
	log.Println(config.Config.Api.QueueName)
	log.Println(config.Config.DataIngestion.QueueName)
	// messaging.PublishMessage()
	messenger := messaging.Messenger{QueueName: config.Config.Api.QueueName}
	messenger.Publish("sending a message from the api service", config.Config.DataIngestion.QueueName)
}
