package repo

import (
	"lets/models"

	"github.com/jinzhu/gorm"
)

type ProductRepo struct {
	DB *gorm.DB `inject:""`
}

func (r *ProductRepo) FindAll() []models.Product {

	var products []models.Product
	r.DB.Limit(10).Find(&products)
	return products
}
