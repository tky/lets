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
}

func (r *ProductRepoImpl) FindAll() []models.Product {

	var products []models.Product
	r.DB.Limit(10).Find(&products)
	return products
}
