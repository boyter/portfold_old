package forms

import (
	"testing"
)

func TestUserHash(t *testing.T) {
	form := NewForm(map[string][]string{})

	form.Required("title")
	form.MaxLength("title", 10)
	if form.Valid() {
		t.Error("Expecting form to be invalid")
	}
}
