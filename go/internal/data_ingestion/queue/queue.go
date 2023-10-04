package queue

import (
	"encoding/json"
	"log"

	"github.com/JWindy92/MovieDataCruncher/config"
	"github.com/JWindy92/MovieDataCruncher/internal/data_ingestion/ingestion"
	"github.com/JWindy92/MovieDataCruncher/messaging"
	"github.com/JWindy92/MovieDataCruncher/models"
)

// TODO: add logic for determining which func to run
func HandleMsg(s string) {
	log.Printf("Ingesting data for %s", s)

	// 1. Check redis for title_year
	data := ingestion.PerformSearch(s)

	redisCacheMsg := models.RedisCacheIdMsg{}
	json.Unmarshal([]byte(data), &redisCacheMsg)
	json_str, err := json.Marshal(redisCacheMsg)
	if err != nil {
		log.Println(err)
	}
	messaging.Publish(string(json_str), config.Config.Redis.QueueName)
	messaging.Publish(string(data), config.Config.Api.QueueName)
}
