package queue

import (
	"log"

	"github.com/JWindy92/MovieDataCruncher/config"
	"github.com/JWindy92/MovieDataCruncher/internal/data_ingestion/ingestion"
	"github.com/JWindy92/MovieDataCruncher/messaging"
)

// TODO: add logic for determining which func to run
func HandleMsg(s string) {
	log.Printf("Ingesting data for %s", s)
	data := ingestion.PerformSearch(s)
	messaging.Publish(string(data), config.Config.Api.QueueName)
}
