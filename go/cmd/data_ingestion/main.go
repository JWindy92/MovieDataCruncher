package main

// This service listens to requests to retrieve information via it's rabbit mq queue
// it is also responsible for caching the unique id's of records as they are read, so other services have more direct access to relevant data
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
