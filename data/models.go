package data

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("data: no matching record found")

type Project struct {
	Id      int
	Name    string
	Created time.Time
	Updated time.Time
}
