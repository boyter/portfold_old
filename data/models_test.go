package data

import (
	"testing"
)

func TestUserHash(t *testing.T) {
	user := User{}
	_ = user.HashPassword("something")

	if user.HashedPassword == nil {
		t.Error("Password should now be hashed")
	}

}
