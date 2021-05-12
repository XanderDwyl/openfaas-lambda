package function

import (
	"os"

	"github.com/XanderDwyl/openfaas-lambda/svc"
)

// Handle a serverless request
func Handle(req []byte) string {
	svc.LogIt("CloudStorage v1.3")

	// Make sure to set GOOGLE_APPLICATION_CREDENTIALS env
	body, err := svc.GCSGetFromKeyWithConfig(os.Getenv("FILE_LOCATION"), os.Getenv("BUCKET_NAME"), false)
	if err != nil {
		return err.Error()
	}

	svc.LogIt(string(body))

	return "CloudStorage"
}
