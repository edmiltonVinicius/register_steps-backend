package cache

import "time"

const (
	TTL_ONE_MINUTE = 1000
	TTL_ONE_DAY    = time.Hour * 24
	TLL_TEN_DAYS   = time.Hour * 24 * 10
)

type Cache struct{}

func NewCache() ICache {
	return &Cache{}
}

type ICache interface {
	CheckStatusConnection() (err error)
	SetJSon(key string, value interface{}, ttl time.Duration) (err error)
	GetJSon(key string, value interface{}) (err error)
}
