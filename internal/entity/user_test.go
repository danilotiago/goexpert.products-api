package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	expectedName := "John Doe"
	expectedEmail := "j@j.com"
	expectedPassword := "123456"

	user, err := NewUser(expectedName, expectedEmail, expectedPassword)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, expectedName, user.Name)
	assert.Equal(t, expectedEmail, user.Email)
}

func TestIsPasswordValid(t *testing.T) {
	expectedName := "John Doe"
	expectedEmail := "j@j.com"
	expectedPassword := "123456"

	user, _ := NewUser(expectedName, expectedEmail, expectedPassword)
	assert.NotNil(t, user)

	assert.True(t, user.IsPasswordValid(expectedPassword))
	assert.False(t, user.IsPasswordValid("123"))
	assert.NotEqual(t, expectedPassword, user.Password)
}
