package function

import (
	"fmt"
	"os"

	svc "github.com/XanderDwyl/openfaas-lambda/svc"
)

// Handle a serverless request
func Handle(req []byte) string {

	snsAttributes:= svc.AttributesKey{}
	err = svc.SNSMessage([]byte("OpenFaas SNS Message"), &snsAttributes)

	return fmt.Sprintf("OpenFaas successfully send")
}
