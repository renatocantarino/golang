package database

import (
	"github.com/google/uuid"
	"github.com/renatocantarino/go/APIS/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{DB: db}

}

func (p *Product) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (prd *Product) FindById(identifier uuid.UUID) (*entity.Product, error) {
	var produt entity.Product
	err := prd.DB.First(&produt, "id = ?", identifier).Error
	return &produt, err
}

func (prd *Product) FindByName(name string) (*entity.Product, error) {
	var produt entity.Product
	err := prd.DB.First(&produt, "name = ?", name).Error
	return &produt, err
}

func getSort(sort string) string {
	if sort != "" && sort != "asc" && sort != "desc" {
		return "asc"
	}

	return sort
}

func (prd *Product) FindAll(page, limit int, sort string) ([]entity.Product, error) {

	var produtos []entity.Product
	var err error
	_sort := getSort(sort)
	if page != 0 && limit != 0 {
		err = prd.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + _sort).Find(&produtos).Error

	} else {
		err = prd.DB.Order("created_at " + sort).Find(&produtos).Error
	}

	return produtos, err

}

func (prd *Product) Update(prod *entity.Product) error {
	updated, err := prd.FindById(prod.ID)
	if err != nil {
		return err
	}

	prod.CreatedAt = updated.CreatedAt

	return prd.DB.Save(prod).Error

}

func (prd *Product) Delete(identifier uuid.UUID) error {
	saved, err := prd.FindById(identifier)
	if err != nil {
		return err
	}

	return prd.DB.Delete(saved).Error
}
