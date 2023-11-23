package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := CreateProduct("Prod1", 10.50)

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "Prod1", p.Name)
	assert.Equal(t, 10.50, p.Price)
}

func TestProductWhenNameIsEmpty(t *testing.T) {
	p, err := CreateProduct("", 10.50)

	assert.Nil(t, p)
	assert.Equal(t, ErrNameIsRequired, err)

}

func TestProductWhenPriceIsEmpty(t *testing.T) {
	p, err := CreateProduct("Prod", 0)

	assert.Nil(t, p)
	assert.Equal(t, ErrPriceIsRequired, err)

}

func TestProductWhenPriceIsLessZero(t *testing.T) {
	p, err := CreateProduct("Prod", -50)

	assert.Nil(t, p)
	assert.Equal(t, ErrPriceIsInvalid, err)

}
