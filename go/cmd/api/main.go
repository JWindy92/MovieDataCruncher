package main

import (
	"log"

	"github.com/JWindy92/MovieDataCruncher/config"
	"github.com/JWindy92/MovieDataCruncher/internal/api/queue"
	"github.com/JWindy92/MovieDataCruncher/internal/api/server"
	"github.com/JWindy92/MovieDataCruncher/messaging"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		panic(err)
	}
	messenger := messaging.Messenger{QueueName: config.Config.Api.QueueName}
	go messenger.Listen(queue.HandleMsg)
	log.Printf("Sending to queue: %s", config.Config.DataIngestion.QueueName)
	messaging.Publish("sending a message from the api service", config.Config.DataIngestion.QueueName)
	server.Run()
}

func testMsg() {
	log.Println(config.Config.Api.QueueName)
	log.Println(config.Config.DataIngestion.QueueName)
	// messenger := messaging.Messenger{QueueName: config.Config.Api.QueueName}
	messaging.Publish("sending a message from the api service", config.Config.DataIngestion.QueueName)
}
