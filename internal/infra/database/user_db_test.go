package database

import (
	"testing"

	entities "github.com/renatocantarino/go/APIS/internal/Entities"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entities.User{})
	user, _ := entities.CreateUser("renato", "r@r.com.br", "123456", "09885984477")
	userDB := NewUser(db)

	err = userDB.Create(user)

	assert.Nil(t, err)

	userfound, err := userDB.FindById(user.ID)

	assert.NotNil(t, userfound)
	assert.NotEmpty(t, userfound.ID)
	assert.NotEmpty(t, userfound.Email)
	assert.NotEmpty(t, userfound.Document)
	assert.Equal(t, "renato", userfound.Name)
	assert.Equal(t, "09885984477", userfound.Document)
	assert.True(t, userfound.ValidatePassword("123456"))
	assert.False(t, userfound.ValidatePassword("12344456"))
	assert.NotEqual(t, "123456", userfound.Password)

}
