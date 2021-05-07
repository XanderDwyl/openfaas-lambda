package function

import (
	"fmt"

	db "github.com/XanderDwyl/openfaas-lambda/svc"
)

// Handle a serverless request
func Handle(req []byte) string {
	db, dbErr := db.GetDBConnection("Sample OpenFaas DB Connection")
	if dbErr != nil {
		db.LogIt(
			fmt.Sprintf(
				"Fetch API Key Error: %s",
				dbErr.Error(),
			),
		)
		return "DB Connection error"
	}

	return fmt.Sprintf("Connection DB: %v", db)
}
