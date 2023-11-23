package database

import (
	"github.com/google/uuid"
	entities "github.com/renatocantarino/go/APIS/internal/Entities"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{DB: db}

}

func (prd *Product) Create(prod *entities.Product) error {
	return prd.DB.Create(prod).Error
}

func (prd *Product) FindById(identifier uuid.UUID) (*entities.Product, error) {
	var produt entities.Product
	err := prd.DB.First(&produt, "id = ?", identifier).Error
	return &produt, err
}

func getSort(sort string) string {
	if sort != "" && sort != "asc" && sort != "desc" {
		return "asc"
	}

	return sort
}

func (prd *Product) FindAll(page, limit int, sort string) ([]entities.Product, error) {

	var produtos []entities.Product
	var err error
	_sort := getSort(sort)
	if page != 0 && limit != 0 {
		err = prd.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + _sort).Find(&produtos).Error

	} else {
		err = prd.DB.Order("created_at " + sort).Find(&produtos).Error
	}

	return produtos, err

}

func (prd *Product) Update(prod *entities.Product) error {
	_, err := prd.FindById(prod.ID)
	if err != nil {
		return err
	}

	return prd.DB.Save(prod).Error

}

func (prd *Product) Delete(identifier uuid.UUID) error {
	saved, err := prd.FindById(identifier)
	if err != nil {
		return err
	}

	return prd.DB.Delete(saved).Error
}
