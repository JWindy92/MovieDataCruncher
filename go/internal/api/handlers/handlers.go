package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/JWindy92/MovieDataCruncher/config"
	"github.com/JWindy92/MovieDataCruncher/messaging"
	"github.com/JWindy92/MovieDataCruncher/models"
	"github.com/gin-gonic/gin"
)

func MovieHandler() gin.HandlerFunc {

	//  - Recieve GET
	//  - Ask data_ingestion to get data
	//	- Listener listens for message from data_ingestion containing data
	//  - Api sends websocket update to client
	return func(c *gin.Context) {
		body := models.SearchQuery{}
		if err := c.BindJSON(&body); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		log.Printf("Asking data ingestion for %s", body.Title)
		body_str, err := json.Marshal(body)
		if err != nil {
			panic(err)
		}
		log.Printf("Json String: %s", body_str)
		messaging.Publish(string(body_str), config.Config.DataIngestion.QueueName)
		c.JSON(http.StatusOK, gin.H{
			"message": body,
		})
	}
}
