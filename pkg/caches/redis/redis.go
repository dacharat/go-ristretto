package redis

import (
	"context"
	"fmt"

	"github.com/dacharat/go-ristretto/pkg/config"
	redisClient "github.com/go-redis/redis/v8"
)

func NewRedis() *redisClient.Client {
	fmt.Println("Redis Host: ", config.Cfg.Redis.Host)
	client := redisClient.NewClient(&redisClient.Options{
		Addr:     config.Cfg.Redis.Host,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}
	return client
}
