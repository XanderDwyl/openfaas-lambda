package function

import (
	"fmt"
	"strconv"
	"time"

	"gopkg.in/redis.v5"

	svc "github.com/XanderDwyl/openfaas-lambda/svc"
)

var redisc *redis.Client

// Handle a serverless request
func Handle(req []byte) string {

	redisc, err := svc.GetRedisConnection()
	if err != nil {
		return fmt.Sprintf("Redis Connection Error: %v", err.Error())
	}

	idRedisKey, _ := strconv.Atoi(string(req))
	if isRedisSetAlready(redisc, idRedisKey) {
		msg := "Redis key is set"
		svc.LogIt(msg)

		keys, err := redisc.Get(string(req)).Result()
		svc.LogIt(fmt.Sprintf("Keys => %v | %v", keys, err))
		return "REDIS is already set"
	}

	// Otherwise, we set it
	err = redisSetCampaign(redisc, idRedisKey)
	if err != nil {
		svc.LogIt(
			fmt.Sprintf(
				"Setting redis with key error : %s",
				err.Error(),
			),
		)
		return "Could not set redis with key"
	}

	return "DONE"
}

func isRedisSetAlready(redisc *redis.Client, cid int) bool {
	key := strconv.Itoa(cid)
	val := redisc.Get(key).Val()

	svc.LogIt(fmt.Sprintf("Redis Value: %v", val))

	return val == key
}

func redisSetCampaign(redisc *redis.Client, cid int) error {
	key := strconv.Itoa(cid)

	// Set expiration to 2 hours
	rediscErr := redisc.Set(key, key, 2*time.Hour).Err()

	return rediscErr
}
