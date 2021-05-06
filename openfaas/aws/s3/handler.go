package function

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"

	utility "github.com/XanderDwyl/openfaas-go-examples"
)

// Handle a serverless request
func Handle(req []byte) string {
	key := os.Getenv("S3_KEY")

	creds := credentials.NewStaticCredentials(os.Getenv("AWS_ID"), os.Getenv("AWS_KEY"), "")
	_, err := creds.Get()
	if err != nil {
		// handle error
	}

	cfg := aws.NewConfig().WithRegion(os.Getenv("REGION")).WithCredentials(creds)

	b, err := utility.S3GetFromKeyWithConfig(key, "wavecell.new.dev.cp", cfg, true)
	if err != nil {
		return fmt.Sprintf("%v: Could not get file from S3", err.Error())
	}

	// Encode to payload
	var payload utility.ImportPayload
	err = json.Unmarshal(b, &payload)
	if err != nil {
		return "Could not unmarshal payload"
	}

	fmt.Printf("Payload: %v", payload)

	return "Success S3 Process"
}
