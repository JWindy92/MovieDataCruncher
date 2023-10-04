package main

import (
	"context"

	"github.com/JWindy92/MovieDataCruncher/config"
	"github.com/JWindy92/MovieDataCruncher/internal/redis_service/cache"
	"github.com/JWindy92/MovieDataCruncher/messaging"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		// Handle error, e.g., log and exit
		panic(err)
	}
	ctx := context.TODO()
	cache.ConnectRedis(ctx)
	messenger := messaging.Messenger{QueueName: config.Config.Redis.QueueName}
	messenger.Listen(cache.HandleQueueMsg)
}
