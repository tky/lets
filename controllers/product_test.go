package controllers

import (
	"encoding/json"
	"lets/controllers"
	"lets/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/facebookgo/inject"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/stretchr/testify/assert"
)

type MockProductRepo struct{}

func (r *MockProductRepo) FindAll() []models.Product {
	return []models.Product{
		models.Product{
			Code:  "ABC",
			Price: 10,
		},
	}

}

func TestIndex(t *testing.T) {
	var productRepo MockProductRepo

	var ctrl controllers.ProductController
	inject.Populate(&ctrl, &productRepo)

	e := echo.New()
	req := new(http.Request)
	rec := httptest.NewRecorder()
	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
	c.SetPath("/")

	if assert.NoError(t, ctrl.Index(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var products []models.Product
		json.Unmarshal(([]byte)(rec.Body.String()), &products)
		assert.Equal(t, 1, len(products))
		assert.Equal(t, "ABC", products[0].Code)
		assert.Equal(t, 10, (int)(products[0].Price))
	}
}
