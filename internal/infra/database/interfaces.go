package database

import entities "github.com/renatocantarino/go/APIS/internal/Entities"

type UserInterface interface {
	Create(user *entities.User) error
	FindByEmail(email string) (*entities.User, error)
}
