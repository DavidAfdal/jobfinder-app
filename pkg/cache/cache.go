package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/DavidAfdal/workfinder/config"
	"github.com/redis/go-redis/v9"
)



func InitCache(config *config.RedisConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Password: config.Password,
		DB:       0,
	})

	return rdb
}

type Cacheable interface {
	Get(key string) string
	Set(key string, value interface{}, expire time.Duration) error
	Delete(key string) error
}

type cacheable struct {
	Redis *redis.Client
}


func NewCacheable(redis *redis.Client) Cacheable {
	return &cacheable{Redis: redis}
}


func (c *cacheable) Set(key string, value interface{}, expire time.Duration) error {
	if err := c.Redis.Set(context.Background(), key, value, expire).Err(); err != nil {
		return err
	}

	return nil
}

func (c *cacheable) Get(key string) string {
	val, err := c.Redis.Get(context.Background(), key).Result()

	if err == redis.Nil {
		return ""
	}

	return val
}

func (c *cacheable) Delete(key string) error {
	_, err := c.Redis.Del(context.Background(), key).Result()

	if err != nil {
		return err
	}

	return nil
}
