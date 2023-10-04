package cache

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client  *redis.Client
	Context context.Context
}

var redis_obj Redis

func ConnectRedis(ctx context.Context) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Println(err)
	}
	log.Println(pong)

	redis_obj = Redis{
		Client:  client,
		Context: ctx,
	}

}

func HandleQueueMsg(msg string) {
	log.Println("Redis service recieved a message")
}

func SetRedis(ctx context.Context, key, val string) {
	err := redis_obj.Client.Set(ctx, key, val, 0).Err()
	if err != nil {
		log.Println(err)
	}
}

func GetRedis(ctx context.Context, key string) string {
	val, err := redis_obj.Client.Get(ctx, key).Result()
	if err != nil {
		log.Println(err)
	}

	return val
}
