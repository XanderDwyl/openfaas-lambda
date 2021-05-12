package svc

import (
	"fmt"
	"os"

	"gopkg.in/redis.v5"
)

func GetRedisConnection(sentinelAddrs []string) (*redis.Client, error) {
	LogIt(fmt.Sprintf("Address: %v", sentinelAddrs))

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
