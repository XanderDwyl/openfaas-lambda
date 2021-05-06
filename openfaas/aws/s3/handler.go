package function

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	svc "github.com/XanderDwyl/openfaas-lambda/svc"
)

// Handle a serverless request
func Handle(req []byte) string {
	cfg := svc.GetConfig()
	s3SQS := s3.New(session.New(), cfg)

	b, err := svc.S3GetFromKeyWithConfig(os.Getenv("S3_KEY"), os.Getenv("S3_BUCKET"), s3SQS, true)
	if err != nil {
		return fmt.Sprintf("%v: Could not get file from S3\n\n", err.Error())
	}

	fmt.Printf("Payload: %v\n\n", b)

	return "Success S3 Process"
}
