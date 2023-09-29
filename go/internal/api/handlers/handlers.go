package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/JWindy92/MovieDataCruncher/config"
	"github.com/JWindy92/MovieDataCruncher/messaging"
	"github.com/gin-gonic/gin"
)

type MovieQuery struct {
	Title string `json:"title"`
}

func MovieHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		body := MovieQuery{}
		if err := c.BindJSON(&body); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		// msgr := messaging.Messenger{QueueName: config.Config.Api.QueueName}
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
