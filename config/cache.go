package config

import "github.com/go-redis/redis/v8"

func OpenCache(redisURL string) (*redis.Client, error) {
	options, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}

	rdb := redis.NewClient(options)

	return rdb, nil
}
