package cache

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/edmiltonVinicius/register-steps/config"
	"github.com/redis/go-redis/v9"
)

func (c *Cache) CheckStatusConnection() (err error) {
	if config.RedisClient == nil {
		err = errors.New("redis no connected")
		return
	}
	err = config.RedisClient.Ping(config.Environment.CTX).Err()
	return
}

func (c *Cache) SetJSon(key string, value interface{}, ttl time.Duration) (err error) {
	err = c.CheckStatusConnection()
	if err != nil {
		return
	}

	res, err := json.Marshal(value)
	if err != nil {
		return
	}
	err = config.RedisClient.Set(config.Environment.CTX, key, res, ttl).Err()
	return
}

func (c *Cache) GetJSon(key string, value interface{}) (err error) {
	err = c.CheckStatusConnection()
	if err != nil {
		return
	}

	res, err := config.RedisClient.Get(config.Environment.CTX, key).Result()
	if err == redis.Nil {
		err = errors.New("key not found")
		return
	}

	if err != nil {
		return
	}

	if res == "" {
		err = errors.New("value key is empty")
		return
	}

	err = json.Unmarshal([]byte(res), value)
	return
}
