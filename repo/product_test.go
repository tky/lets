package repo_test

import (
	"lets/models"
	"lets/repo"
	"os"
	"testing"

	"github.com/facebookgo/inject"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func MakeTestDB() *gorm.DB {
	// I want to delete test.db after executing tests.
	os.Remove("test.db")

	DB, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&models.Product{})

	return DB
}

func TestFindAll(t *testing.T) {
	var repo repo.ProductRepoImpl
	db := MakeTestDB()
	db.Create(&models.Product{Code: "1234", Price: 1000})
	defer db.Close()

	inject.Populate(&repo, db)

	products := repo.FindAll()
	assert.Equal(t, 1, len(products))
}

func TestFind(t *testing.T) {
	var repo repo.ProductRepoImpl
	db := MakeTestDB()
	db.Create(&models.Product{Code: "1234", Price: 1000})
	defer db.Close()

	inject.Populate(&repo, db)

	id := repo.FindAll()[0].Id

	product, err := repo.Find(id)
	assert.Nil(t, err)
	assert.Equal(t, product.Id, id)
	assert.Equal(t, product.Code, "1234")
	assert.Equal(t, product.Price, 1000)
}

func TestFindIfNotFound(t *testing.T) {
	var repo repo.ProductRepoImpl
	db := MakeTestDB()
	db.Create(&models.Product{Code: "1234", Price: 1000})
	defer db.Close()

	inject.Populate(&repo, db)

	_, err := repo.Find(99999)

	assert.NotNil(t, err)
	assert.Equal(t, "not found [id=99999]", err.Error())
}

func TestSaveAndFind(t *testing.T) {
	var repo repo.ProductRepoImpl
	db := MakeTestDB()
	defer db.Close()

	inject.Populate(&repo, db)

	if product, err := repo.Save(&models.Product{
		Code:  "A01",
		Price: 100,
	}); err != nil {
		t.Error("Should not return error")
	} else {
		assert.NotNil(t, product.Id, "Should set Id")

		found, _ := repo.Find(product.Id)
		assert.Equal(t, "A01", found.Code)
		assert.Equal(t, 100, found.Price)
	}

}
