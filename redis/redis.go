package cache

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/edmiltonVinicius/register-steps/domain"
	"github.com/redis/go-redis/v9"
)

const (
	TTL_ONE_MINUTE = 1000
	TTL_ONE_DAY    = time.Hour * 24
	TLL_TEN_DAYS   = time.Hour * 24 * 10
)

func checkStatusConnection() (err error) {
	err = domain.RedisClient.Ping(domain.Environment.CTX).Err()
	return
}

func SetJSon(key string, value interface{}, ttl time.Duration) (err error) {
	err = checkStatusConnection()
	if err != nil {
		return
	}

	res, err := json.Marshal(value)
	if err != nil {
		return
	}
	err = domain.RedisClient.Set(domain.Environment.CTX, key, res, ttl).Err()
	return
}

func GetJSon(key string, value interface{}) (err error) {
	err = checkStatusConnection()
	if err != nil {
		return
	}

	res, err := domain.RedisClient.Get(domain.Environment.CTX, key).Result()
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
