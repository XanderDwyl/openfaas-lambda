package svc

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

func GetSNSConfig() *aws.Config {
	creds := credentials.NewStaticCredentials(os.Getenv("AWS_ID"), os.Getenv("AWS_KEY"), "")
	_, err := creds.Get()
	if err != nil {
		fmt.Printf("Error: %v", err.Error())
		return nil
	}

	return aws.NewConfig().WithRegion(os.Getenv("REGION")).WithCredentials(creds)

}
