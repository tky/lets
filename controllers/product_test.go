package controllers_test

/*
import (
	"encoding/json"
	"lets/controllers"
	"lets/models"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/echo-contrib/pongor"
	"github.com/facebookgo/inject"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/stretchr/testify/assert"
)

type MockProductRepo struct {
	Saved *models.Product
}

func (r *MockProductRepo) FindAll() []models.Product {
	return []models.Product{
		models.Product{
			Code:  "ABC",
			Price: 10,
		},
	}
}

func (r *MockProductRepo) Find(id int) (*models.Product, error) {
	panic("")
}

func (r *MockProductRepo) Save(m *models.Product) (*models.Product, []error) {
	r.Saved = m
	return nil, nil
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

func TestPost(t *testing.T) {
	var productRepo MockProductRepo

	var ctrl controllers.ProductController
	inject.Populate(&ctrl, &productRepo)

	f := make(url.Values)
	f.Set("Code", "B1")
	f.Set("Price", "200")

	e := echo.New()
	req, _ := http.NewRequest(echo.POST, "/", strings.NewReader(f.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()
	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
	c.SetPath("/")

	if assert.NoError(t, ctrl.Post(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "B1", productRepo.Saved.Code)
	} else {
		t.Error("Should not return error")
	}
}

func setTemplates(e *echo.Echo) {
	r := pongor.GetRenderer(pongor.PongorOption{
		Directory: "../views/",
	})
	e.SetRenderer(r)
}

func TestIndexHtml(t *testing.T) {
	var productRepo MockProductRepo

	var ctrl controllers.ProductController
	inject.Populate(&ctrl, &productRepo)

	e := echo.New()
	setTemplates(e)
	req := new(http.Request)
	rec := httptest.NewRecorder()
	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
	c.SetPath("/product.html")

	if assert.NoError(t, ctrl.IndexHtml(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		doc, _ := goquery.NewDocumentFromReader(strings.NewReader(rec.Body.String()))
		assert.Equal(t, doc.Find("#list").Find("li").Text(), "ABC : 10")
	}
}
*/
