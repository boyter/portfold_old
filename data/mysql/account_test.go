// +build integration

package mysql

import (
	"testing"
)

func TestAccountGet(t *testing.T) {
	accountModel := AccountModel{DB: connect(t)}
	_, err := accountModel.GetAccount(2147483647)

	if err.Error() != "data: no matching record found" {
		t.Error("Expected to get no matching record")
	}
}

//func TestAccountInsert(t *testing.T) {
//	accountModel := AccountModel{DB: connect(t)}
//	account, err := accountModel.Insert(data.Account{
//		Name:    "test account",
//		Email:   "test@example.com",
//		Details: "some details about this account",
//	})
//
//	if err != nil {
//		t.Error(err.Error())
//	}
//
//	accountModel.GetAccount(account.Id)
//	accountModel.Delete(*account)
//	_, err = accountModel.GetAccount(account.Id)
//
//	if err.Error() != "data: no matching record found" {
//		t.Error("Expected to get no matching record")
//	}
//}
