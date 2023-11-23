package entities

import (
	"errors"
	"time"

	entities "github.com/renatocantarino/go/APIS/pkg/Entities"
)

var (
	ErrIdIsRequired    = errors.New("Id is Required")
	ErrNameIsRequired  = errors.New("Name is Required")
	ErrPriceIsRequired = errors.New("Price is Required")
	ErrPriceIsInvalid  = errors.New("Price is Invalid")
)

type Product struct {
	ID       entities.ID `json:"id"`
	Name     string      `json:"name"`
	Price    float64     `json:"price"`
	CreateAt time.Time   `json:"created_at"`
}

func CreateProduct(name string, price float64) (*Product, error) {

	product := &Product{
		ID:       entities.Generate(),
		Name:     name,
		Price:    price,
		CreateAt: time.Now(),
	}

	err := product.Validate()
	if err != nil {
		return nil, err
	}

	return product, nil

}

func (p *Product) Validate() error {

	if p.ID.String() == "" {
		return ErrIdIsRequired
	}

	if _, err := entities.Validate(p.ID.String()); err != nil {
		return ErrIdIsRequired
	}

	if p.Name == "" {
		return ErrNameIsRequired
	}

	if p.Price < 0 {
		return ErrPriceIsInvalid
	}

	if p.Price == 0 {
		return ErrPriceIsRequired
	}

	return nil

}
