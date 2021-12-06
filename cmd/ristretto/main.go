package main

import (
	"github.com/dacharat/go-ristretto/cmd/ristretto/route"
	"github.com/dacharat/go-ristretto/pkg/caches"
	"github.com/dacharat/go-ristretto/pkg/caches/redis"
	"github.com/dacharat/go-ristretto/pkg/caches/ristretto"
	"github.com/dacharat/go-ristretto/pkg/config"
)

func main() {
	config.Process()
	redisClient := redis.NewRedis()
	localCache := ristretto.NewLocalcahe()
	c := caches.NewCache(redisClient, localCache)

	router := route.NewRouter(c)
	router.Run()
}
