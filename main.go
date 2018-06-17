package main

import (
	"log"
	"time"
)

func main() {
	opt := RedisOption{
		Host:     "localhost",
		Port:     "6379",
		Password: "",
		Database: 0,
	}
	redis := NewRedis(opt)
	redis.db.Set("key", "value", 10*time.Hour)
	log.Println(redis.db.Get("key"))

	log.Println(redis.db.Get("asu"))
}
