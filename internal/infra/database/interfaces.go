package database

import (
	"github.com/google/uuid"
	entities "github.com/renatocantarino/go/APIS/internal/Entities"
)

type UserInterface interface {
	Create(user *entities.User) error
	FindByEmail(email string) (*entities.User, error)
	FindById(identifier uuid.UUID) (*entities.User, error)
}

type ProductInterface interface {
	Create(prod *entities.Product) error
	FindAll(page, limit int, sort string) ([]entities.Product, error)
	FindByName(name string) (entities.Product, error)
	FindById(identifier uuid.UUID) (*entities.Product, error)
	Update(prd *entities.Product) error
	Delete(identifier uuid.UUID) error
}
