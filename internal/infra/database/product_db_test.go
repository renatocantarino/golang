package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/renatocantarino/go/APIS/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func createDbConex() (*gorm.DB, error) {

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db, nil
}

func TestNewProduct(t *testing.T) {

	db, err := createDbConex()
	db.AutoMigrate(&entity.Product{})

	productDB := NewProduct(db)

	productCreated, _ := entity.CreateProduct("produto1", 150.55)
	err = productDB.Create(productCreated)

	assert.Nil(t, err)

	prdfound, err := productDB.FindById(productCreated.ID)

	assert.NotNil(t, prdfound)
	assert.NotEmpty(t, prdfound.ID)
	assert.Equal(t, "produto1", prdfound.Name)
	assert.Equal(t, 150.55, prdfound.Price)

}

func TestProductsPagination(t *testing.T) {
	db, err := createDbConex()

	db.AutoMigrate(&entity.Product{})

	for i := 1; i < 30; i++ {
		productCreated, err := entity.CreateProduct(fmt.Sprintf("Produto-%d", i), rand.Float64()*100)
		assert.NoError(t, err)
		db.Create(productCreated)
	}

	productDB := NewProduct(db)
	all, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, all, 10)
	assert.Equal(t, "Produto-1", all[0].Name)
	assert.Equal(t, "Produto-10", all[9].Name)

	all, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, all, 10)
	assert.Equal(t, "Produto-11", all[0].Name)
	assert.Equal(t, "Produto-20", all[9].Name)

}

func TestFindProductById(t *testing.T) {
	db, err := createDbConex()
	db.AutoMigrate(&entity.Product{})
	productCreated, _ := entity.CreateProduct("produto1", 150.55)
	productDB := NewProduct(db)

	err = productDB.Create(productCreated)

	assert.Nil(t, err)

	prdfound, err := productDB.FindById(productCreated.ID)

	assert.NotNil(t, prdfound)
	assert.NotEmpty(t, prdfound.ID)
	assert.Equal(t, "produto1", prdfound.Name)
	assert.Equal(t, 150.55, prdfound.Price)
}

func TestFindProductByName(t *testing.T) {
	db, err := createDbConex()
	db.AutoMigrate(&entity.Product{})
	productCreated, _ := entity.CreateProduct("produto10", 150.55)
	productDB := NewProduct(db)

	err = productDB.Create(productCreated)

	assert.Nil(t, err)

	prdfound, err := productDB.FindByName("produto10")

	assert.NotNil(t, prdfound)
	assert.NotEmpty(t, prdfound.ID)
	assert.Equal(t, "produto10", prdfound.Name)
	assert.Equal(t, 150.55, prdfound.Price)
}

func TestDeleteProduct(t *testing.T) {
	db, err := createDbConex()

	db.AutoMigrate(&entity.Product{})
	productCreated, _ := entity.CreateProduct("produto10", 150.55)
	productDB := NewProduct(db)

	err = productDB.Create(productCreated)

	assert.Nil(t, err)

	prdfound, err := productDB.FindByName("produto10")

	assert.NotNil(t, prdfound)
	assert.NotEmpty(t, prdfound.ID)
	assert.Equal(t, "produto10", prdfound.Name)
	assert.Equal(t, 150.55, prdfound.Price)

	err = productDB.Delete(prdfound.ID)

	_, err2 := productDB.FindByName("produto10")

	assert.Error(t, err2)

}
