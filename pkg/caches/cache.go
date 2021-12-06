package caches

import (
	"context"
	"fmt"
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/go-redis/redis/v8"
)

type ICache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
}

type Cache struct {
	redis *redis.Client
	local *ristretto.Cache
}

func NewCache(redis *redis.Client, local *ristretto.Cache) ICache {
	return &Cache{
		redis: redis,
		local: local,
	}
}

func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	value, exist := c.local.Get(key)
	if exist {
		fmt.Println("Get from 'ristretto'")
		return fmt.Sprintf("%v", value), nil
	}
	result, err := c.redis.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	ttl, err := c.redis.TTL(ctx, key).Result()
	if err != nil {
		fmt.Println("Cache.Get Err: ", err)
	}

	c.local.SetWithTTL(key, result, 2, ttl)
	c.local.Wait()

	fmt.Println("Get from 'redis'")
	return result, nil
}

func (c *Cache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	err := c.redis.Set(ctx, key, value, ttl).Err()
	if err != nil {
		return err
	}

	c.local.SetWithTTL(key, value, 2, ttl)
	c.local.Wait()
	return nil
}
