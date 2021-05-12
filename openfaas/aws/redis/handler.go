package function

import (
	"fmt"

	"gopkg.in/redis.v5"

	svc "github.com/XanderDwyl/openfaas-lambda/svc"
)

var redisc *redis.Client

// Handle a serverless request
func Handle(req []byte) string {

	// // if you create secrets for you confidential ENV CONFIG
	// sent1, err := svc.GetAPISecret("SENTINEL_1")
	// sent2, err := svc.GetAPISecret("SENTINEL_2")
	// sent3, err := svc.GetAPISecret("SENTINEL_3")

	// os.Setenv("SENTINEL_1", sent1)
	// os.Setenv("SENTINEL_2", sent2)
	// os.Setenv("SENTINEL_3", sent3)

	redisc, err := svc.GetRedisConnection()
	if err != nil {
		return fmt.Sprintf("Redis Connection Error: %v", err.Error())
	}

	return fmt.Sprintf("connecting to redis : %v", redisc)

}
