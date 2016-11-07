package repo

import (
	"fmt"
	"lets/models"

	"github.com/jinzhu/gorm"
)

type NotFoundError struct {
	Id int
}

func (err *NotFoundError) Error() string {
	return fmt.Sprintf("not found [id=%d]", err.Id)
}

type ProductRepoImpl struct {
	DB *gorm.DB `inject:""`
}

type ProductRepo interface {
	FindAll() []models.Product
	Find(id int) (*models.Product, error)
	Save(*models.Product) (*models.Product, []error)
}

func (r *ProductRepoImpl) FindAll() []models.Product {
	var products []models.Product
	r.DB.Limit(10).Find(&products)
	return products
}

func (r *ProductRepoImpl) Find(id int) (*models.Product, error) {
	var product models.Product
	if r.DB.Find(&product, id).RecordNotFound() {
		return nil, &NotFoundError{Id: id}
	} else {
		return &product, nil
	}
}

func (r *ProductRepoImpl) Save(m *models.Product) (*models.Product, []error) {
	return m, r.DB.Save(m).GetErrors()
}
