package function

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"

	svc "github.com/XanderDwyl/openfaas-lambda/svc"
)

// Handle a serverless request
func Handle(req []byte) string {

	cfg := svc.GetSNSConfig()
	svc := sqs.New(session.New(), cfg)

	// SQS endpoint
	params := sqs.SQSInput("OpenFaas SQS Connection", os.Getenv("SQS_URL"), "SocketId", "socketID")

	svcMsg, err := svc.SendMessage(params)
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("SNS message : %v", svcMsg)
}
