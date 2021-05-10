package function

import (
	"fmt"

	svc "github.com/XanderDwyl/openfaas-lambda/svc"
)

// Handle a serverless request
func Handle(req []byte) string {
	env, err := svc.GetAPISecret(string(req))
	if err != nil {
		env = nil
	}

	fmt.Printf("%s: %s\n\n", string(req), string(env))

	return fmt.Sprintf("Hello, Go. You said: %s", string(req))
}
