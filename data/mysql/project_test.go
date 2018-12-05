// +build integration

package mysql

import (
	"boyter/portfold/data"
	"database/sql"
	"testing"
)

func db(t *testing.T) *sql.DB {
	db, err := data.ConnectDb()
	if err != nil {
		t.Error("Unable to connect to DB", err.Error())
	}

	return db
}

func TestGet(t *testing.T) {
	project := ProjectModel{DB: db(t)}
	project.Get(1)
}
