package user

import "testing"

func TestNew(t *testing.T) { // HL
	user, err := NewUser("", "password")

	if err == nil {
		t.Fail("expected an error when the username is empty")
	}
	if user != User{} {
		t.Fail("expected an empty user when the username is empty")
	}
}
