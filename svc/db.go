package svc

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

var workingDBURL string

// GetDBConnection ...
func GetDBConnection(title string) (db *sql.DB, err error) {

	endpoints := os.Getenv("DB_URL")
	if endpoints == "" {
		err = errors.New("database url not found")
		return
	}

	dbs := strings.Split(endpoints, "|")
	if len(dbs) == 0 {
		err = errors.New("database url not found")
		return
	}

	// Init DB connection with max number of retries
	t := 0
	for {
		// Use the workingDBURL if it's not empty
		dbu := workingDBURL
		if dbu == "" {
			dbu = dbs[t%len(dbs)]
		}

		db, err = sql.Open("sqlserver", dbu)
		if err != nil {
			LogInfo(fmt.Sprintf("Connection Error: %v", err))
			break
		}

		err = db.Ping()
		if err == nil {
			workingDBURL = dbu
			LogInfo("DB Connected")
			break
		}

		// Set to empty if we can no longer connect to it
		workingDBURL = ""
		LogWarn(fmt.Sprintf(
			"[ %v ] %s: Could not connect to the database [ %s ], retrying... %d",
			err,
			title,
			dbu,
			t,
		))

		time.Sleep(time.Duration(t) * time.Second)
		t++

		// Has reached max retries
		// if t == len(dbs) {
		if (len(dbs) > 1 && t == len(dbs)) || (len(dbs) <= 1 && t == 3) {
			errStr := fmt.Sprintf(
				"Could not connect to the database (Max retries reached) : %d",
				t,
			)
			err = errors.New(errStr)
			LogErr(err.Error())
			break
		}
	}

	return db, err
}

// CloseDBConnection ...
func CloseDBConnection(db *sql.DB) {
	_ = db.Close()

	LogInfo("DB connection closed!")
}
