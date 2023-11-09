package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCreateANewUser(t *testing.T) {
	user, err := NewUser("test@test.com", "password123")
	assert.Nil(t, err)
	assert.Equal(t, "test@test.com", user.Email)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.NotEqual(t, user.Password, "password123")
}

func TestShouldNotCreateANewUserIfEmailIsEmpty(t *testing.T) {
	user, err := NewUser("", "password123")
	assert.Nil(t, user)
	assert.Error(t, err)
}

func TestShouldValidatePassword(t *testing.T) {
	user, err := NewUser("test@test.com", "password123")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("password123"))
	assert.False(t, user.ValidatePassword("1233"))
	assert.NotEqual(t, user.Password, "password123")
}
