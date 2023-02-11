package config

import "github.com/go-redis/redis"

var RedisClient *redis.Client

func ConnectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
