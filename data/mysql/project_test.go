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

func TestGet(t *testing.T) {
	project := ProjectModel{DB: connect(t)}
	_, err := project.Get(2147483647)

	if err.Error() != "data: no matching record found" {
		t.Error("Expected to get no matching record")
	}
}
