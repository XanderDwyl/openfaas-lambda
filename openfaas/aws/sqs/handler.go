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

	cfg := svc.GetConfig()
	svcSQS := sqs.New(session.New(), cfg)

	// SQS endpoint
	params := svc.SQSInput("OpenFaas SQS Connection", os.Getenv("SQS_URL"), "SocketId", "socketID")

	svcMsg, err := svcSQS.SendMessage(params)
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("SNS message : %v", svcMsg)
}
