package function

import (
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"

	svc "github.com/XanderDwyl/openfaas-lambda/svc"
)

// Handle a serverless request
func Handle(req []byte) string {
	_, dbErr := svc.GetDBConnection("Sample OpenFaas DB Connection")
	if dbErr != nil {
		svc.LogIt(
			fmt.Sprintf(
				"Fetch API Key Error: %s",
				dbErr.Error(),
			),
		)
		return "DB Connection error"
	}

	return "Success"
}
