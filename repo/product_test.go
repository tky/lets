package repo

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

	product := repo.Find(id)
	assert.Equal(t, product.Id, id)
	assert.Equal(t, product.Code, "1234")
	assert.Equal(t, product.Price, 1000)
}
