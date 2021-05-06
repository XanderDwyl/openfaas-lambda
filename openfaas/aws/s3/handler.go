package function

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"

	svc "github.com/XanderDwyl/openfaas-lambda/svc"
)

// Handle a serverless request
func Handle(req []byte) string {
	key := os.Getenv("S3_KEY")
	bucket := os.Getenv("S3_BUCKET")

	creds := credentials.NewStaticCredentials(os.Getenv("AWS_ID"), os.Getenv("AWS_KEY"), "")
	_, err := creds.Get()
	if err != nil {
		return err.Error()
	}

	cfg := aws.NewConfig().WithRegion(os.Getenv("REGION")).WithCredentials(creds)

	b, err := svc.S3GetFromKeyWithConfig(key, bucket, cfg, true)
	if err != nil {
		return fmt.Sprintf("%v: Could not get file from S3\n\n", err.Error())
	}

	fmt.Printf("Payload: %v\n\n", b)

	return "Success S3 Process"
}
