// +build integration

package mysql

import (
	"boyter/portfold/data"
	"database/sql"
	"testing"
)

// For Model integration tests we need to be able to connect to the
// database but also want to avoid using all the connections
// so we have this to share it and ensure that we only connect once

var db *sql.DB = nil

// NB this method is not thread safe
func connect(t *testing.T) *sql.DB {
	if db != nil {
		return db
	}

	db, err := data.ConnectDb()
	if err != nil {
		t.Error("Unable to connect to DB", err.Error())
	}

	return db
}
