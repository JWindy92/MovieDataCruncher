package main

import (
	"log"

	"github.com/JWindy92/MovieDataCruncher/config"
	"github.com/JWindy92/MovieDataCruncher/internal/data_ingestion/queue"
	"github.com/JWindy92/MovieDataCruncher/messaging"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		// Handle error, e.g., log and exit
		panic(err)
	}
	log.Println("Starting Data Ingestion Service")
	// messaging.PublishMessage()
	messenger := messaging.Messenger{QueueName: config.Config.DataIngestion.QueueName}
	messenger.Listen(queue.HandleMsg)
}
