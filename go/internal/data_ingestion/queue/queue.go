package queue

import (
	"log"

	"github.com/JWindy92/MovieDataCruncher/internal/data_ingestion/ingestion"
)

// TODO: add logic for determining which func to run
func HandleMsg(s string) {
	log.Printf("Ingesting data for %s", s)
	ingestion.PerformSearch(s)
}
