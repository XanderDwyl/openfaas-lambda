package function

import (
	"encoding/json"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"

	svc "github.com/XanderDwyl/openfaas-lambda/svc"
)

// Handle a serverless request
func Handle(req []byte) string {

	cfg := svc.GetSNSConfig()
	svc := sqs.New(session.New(), cfg)

	var event svc.MsgPayload{
		Msg: "OpenFaas SQS Connection"
	}
	
	// JSON encode payload
	msg, err := json.MarshalIndent(evt, "", "  ")
	if err != nil {
		return err.Error()
	}

	// SQS endpoint
	params := sqs.SQSInput(msg, os.Getenv("SQS_URL"), "SocketId", "socketID")

	svcMsg, err := svc.SendMessage(params)
	if err != nil {
		return err.Error()
	}

}
