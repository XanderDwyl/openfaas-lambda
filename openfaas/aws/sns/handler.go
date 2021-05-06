package function

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"

	svc "github.com/XanderDwyl/openfaas-lambda/svc"
)

// Handle a serverless request
func Handle(req []byte) string {

	cfg := svc.GetConfig()
	svc := sns.New(session.New(), cfg)

	topicArn := os.Getenv("SNS_TOPIC")

	input := &sns.PublishInput{
		Message:  aws.String("OpenFaas SNS Message"),
		TopicArn: aws.String(topicArn),
	}

	snsMsg, err := svc.Publish(input)
	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("OpenFaas: %v", snsMsg)
}
