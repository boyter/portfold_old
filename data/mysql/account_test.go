// +build integration

package mysql

import (
	"testing"
)

func TestAccountGet(t *testing.T) {
	project := AccountModel{DB: connect(t)}
	_, err := project.GetAccount(2147483647)

	if err.Error() != "data: no matching record found" {
		t.Error("Expected to get no matching record")
	}
}
