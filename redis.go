package main

import (
	"log"

	"github.com/go-redis/redis"
)

type Redis struct {
	db *redis.Client
}

// Option holds all necessary options for Redis
type RedisOption struct {
	Host     string
	Port     string
	Password string
	Database int
}

func NewRedis(opt RedisOption) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     opt.Host + ":" + opt.Port,
		Password: opt.Password,
		DB:       opt.Database,
	})
	pong, err := client.Ping().Result()
	log.Println(pong, err)
	return &Redis{db: client}
}
