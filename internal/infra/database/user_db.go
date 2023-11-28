package database

import (
	"github.com/google/uuid"
	"github.com/renatocantarino/go/APIS/internal/entity"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{DB: db}
}

func (usr *User) Create(user *entity.User) error {
	return usr.DB.Create(user).Error
}

func (usr *User) FindByEmail(email string) (*entity.User, error) {
	var usuario entity.User
	if err := usr.DB.Where("email = ?", email).First(&usuario).Error; err != nil {
		return nil, err
	}

	return &usuario, nil
}

func (usr *User) FindById(identifier uuid.UUID) (*entity.User, error) {
	var usuario entity.User
	if err := usr.DB.Where("id = ?", identifier).First(&usuario).Error; err != nil {
		return nil, err
	}

	return &usuario, nil
}
