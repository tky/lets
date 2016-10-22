package repo

import (
	"lets/models"

	"github.com/jinzhu/gorm"
)

type ProductRepoImpl struct {
	DB *gorm.DB `inject:""`
}

type ProductRepo interface {
	FindAll() []models.Product
	Find(id int) models.Product
}

func (r *ProductRepoImpl) FindAll() []models.Product {
	var products []models.Product
	r.DB.Limit(10).Find(&products)
	return products
}

func (r *ProductRepoImpl) Find(id int) models.Product {
	var product models.Product
	r.DB.Find(&product, id)
	return product
}
