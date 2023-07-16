package cachelayer

import (
	"errors"

	"github.com/redis/go-redis/v9"
)

func CacheConfig() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	if rdb == nil {
		return nil, errors.New("could not connect to redis")
	}
	return rdb, nil
}
