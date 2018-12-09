// +build integration

package mysql

import (
	"boyter/portfold/data"
	"testing"
)

func TestAccountGet(t *testing.T) {
	accountModel := AccountModel{DB: connect(t)}
	_, err := accountModel.Get(2147483647)

	if err.Error() != "data: no matching record found" {
		t.Error("Expected to get no matching record")
	}
}

func TestAccountInsertDelete(t *testing.T) {
	accountModel := AccountModel{DB: connect(t)}
	account, err := accountModel.Insert(data.Account{
		Name:    "test account",
		Email:   "test@example.com",
		Details: "some details about this account",
	})

	if err != nil {
		t.Error(err.Error())
	}

	accountModel.Delete(*account)
	_, err = accountModel.Get(account.Id)

	if err.Error() != "data: no matching record found" {
		t.Error("Expected to get no matching record")
	}
}
