// +build integration

package mysql

import (
	"boyter/portfold/data"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestUserGet(t *testing.T) {
	userModel := UserModel{DB: connect(t)}
	_, err := userModel.Get(2147483647)

	if err.Error() != "data: no matching record found" {
		t.Error("Expected to get no matching record", err.Error())
	}
}

func TestUserInsertDelete(t *testing.T) {
	accountModel := AccountModel{DB: connect(t)}
	account, err := accountModel.Insert(data.Account{
		Name:    "test account",
		Email:   "test@example.com",
		Details: "some details about this account",
	})

	if err != nil {
		t.Error(err.Error())
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), 12)

	userModel := UserModel{DB: connect(t)}
	user, err := userModel.Insert(data.User{
		AccountId:      account.Id,
		Name:           "test user",
		Email:          "test@example.com",
		HashedPassword: hashedPassword,
	})

	if err != nil {
		t.Error(err.Error())
	}

	userModel.Delete(*user)
	accountModel.Delete(*account)
}
