package entities

import (
	entities "github.com/renatocantarino/go/APIS/pkg/Entities"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entities.ID `json:"id"`
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	Document string      `json:"document"`
	Password string      `json:"-"`
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func CreateUser(name, email, password, document string) (*User, error) {

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
