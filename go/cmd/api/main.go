package main

import (
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
	server.Run()
}
