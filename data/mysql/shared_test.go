// +build integration

package mysql

import (
	"boyter/portfold/data"
	"database/sql"
	"testing"
)

var db *sql.DB = nil

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
