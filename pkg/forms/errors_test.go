package forms

import (
	"testing"
)

func TestAdd(t *testing.T) {

	err := make(errors)
	err.Add("firstName", "This should produce an error")
	if len(err["firstName"]) != 1 {
		t.Errorf("it should produce what you want")
	}
	if err["firstName"][0] != "This should produce an error" {
		t.Errorf("incorrect error message for 'username' field: got %v want %v",
			err["firstName"][0], "Username is required")
	}
}
func TestGet(t *testing.T) {
	err := make(errors)
	err.Add("firstName", "name")
	value := err.Get("firstName")
	unknownError := err.Get("email")
	if value != "name" {
		t.Error("Big error")
	}
	if unknownError != "" {
		t.Errorf("unexpected error message for 'email' field: got %v want %v",
			unknownError, "")
	}
}
