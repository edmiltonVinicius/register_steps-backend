package domain

import (
	"log"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func StartRedis() {
	_, err := strconv.Atoi(Environment.REDIS_DB)
	if err != nil {
		log.Println("Error in parsing redis_db from string to int: ", err)
		return
	}

	client := redis.NewClient(&redis.Options{
		Addr:     Environment.REDIS_URL,
		Password: Environment.REDIS_PASSWORD,
		DB:       0,
	})

	_, err = client.Ping(Environment.CTX).Result()
	if err != nil {
		log.Println("Error connecting redis: ", err)
		return
	}

	RedisClient = client
	log.Println("CONNECTED WITH REDIS")
}
