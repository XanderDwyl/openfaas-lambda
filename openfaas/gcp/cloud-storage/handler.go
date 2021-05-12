package function

import (
	"os"

	"github.com/XanderDwyl/openfaas-lambda/svc"
)

// Handle a serverless request
func Handle(req []byte) string {
	svc.LogIt("CloudStorage v1.3")

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/var/openfaas/secrets/key.json")
	body, err := svc.GCSGetFromKeyWithConfig("PRAKSMOL-4yY8D/A7D8E764-C2F1-45BF-932A-0C0F826D4070/2021/5/12/4-27-22/websender/test_contact.csv", "wavecell-cp-dev", false)
	if err != nil {
		return err.Error()
	}

	svc.LogIt(string(body))

	return "CloudStorage"
}
