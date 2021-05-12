package svc

import (
	"os"

	"gopkg.in/redis.v5"
)

func GetRedisConnection() (*redis.Client, error) {
	// Default or staging credentials
	sentinelAddrs := []string{
		os.Getenv("SENTINEL_1"),
		os.Getenv("SENTINEL_2"),
		os.Getenv("SENTINEL_3"),
	}

	// Connect to our sentinel servers
	client := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "redismaster",
		SentinelAddrs: sentinelAddrs,
		DB:            1,
		Password:      os.Getenv("SENTINEL_PASS"),
		MaxRetries:    3,
	})

	// Ping!
	err := client.Ping().Err()

	if err != nil {
		return nil, err
	}

	return client, nil
}
