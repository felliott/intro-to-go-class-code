package user

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert" // HL
)

func TestNew(t *testing.T) {
	user, err := NewUser("", "password")

	assert.NotEqual(t, err, nil, "got an error when the username is empty") // HL
	assert.Equal(                                                           // HL
		t,
		err,
		errors.New("The username must be a non-empty string."),
		"got the expected error for an empty username",
	)
	assert.Equal(t, user, User{}, "user returned on error is empty")
}
