package function

import (
	"fmt"

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

	return fmt.Sprintf("connecting to redis : %v", redisc)
}
