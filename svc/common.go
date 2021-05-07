package svc

import (
	"fmt"
	"os"
	"time"
)

func LogIt(message string) {
	logMessage := fmt.Sprintf("%s - %s", time.Now().Format(time.RFC3339), message)
	fmt.Fprintln(os.Stderr, logMessage)
}
