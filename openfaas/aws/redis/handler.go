package function

import (
	"fmt"
	"os"

	"gopkg.in/redis.v5"
)

var redisc *redis.Client

// Handle a serverless request
func Handle(req []byte) string {
	if redisc == nil {
		redisc, err := GetRedisConnection()
		if err != nil {
			fmt.Printf("Redis Connection Error: %v", err.Error())

			return "Could not set redis connection hello"
		}

		fmt.Printf("connecting to redis : %v", redisc)
	}

	return "================ Hello Redisc ================"
}

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
