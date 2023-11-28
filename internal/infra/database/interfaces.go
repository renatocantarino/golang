package database

import (
	"github.com/google/uuid"
	"github.com/renatocantarino/go/APIS/internal/entity"
)

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	FindById(identifier uuid.UUID) (*entity.User, error)
}

type ProductInterface interface {
	Create(product *entity.Product) error
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	FindById(identifier uuid.UUID) (*entity.Product, error)
	Update(prd *entity.Product) error
	Delete(identifier uuid.UUID) error
}
