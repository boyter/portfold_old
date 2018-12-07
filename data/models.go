package data

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("data: no matching record found")
var ErrDuplicateEmail = errors.New("data: duplicate email")

type Project struct {
	Id      int
	Name    string
	Created time.Time
	Updated time.Time
}

type Account struct {
	Id      int
	Name    string
	Email   string
	Created time.Time
	Updated time.Time
	Active  bool
	Details string
}
