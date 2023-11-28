package entity

import (
	"github.com/renatocantarino/go/APIS/pkg/entity"
	entities "github.com/renatocantarino/go/APIS/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Document string    `json:"document"`
	Password string    `json:"-"`
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func CreateUser(name, email, document, password string) (*User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return &User{
		ID:       entities.Generate(),
		Name:     name,
		Email:    email,
		Document: document,
		Password: string(hash),
	}, nil
}
