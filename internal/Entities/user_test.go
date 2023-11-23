package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {

	user, err := CreateUser("renato", "r@r.com.br", "123456", "01178844477")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Email)
	assert.NotEmpty(t, user.Document)
	assert.Equal(t, "renato", user.Name)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := CreateUser("renato", "r@r.com.br", "123456", "01178844477")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("12344456"))
	assert.NotEqual(t, "123456", user.Password)

}
