package controllers_test

// mockgen -source repo/product.go

import (
	"encoding/json"
	"lets/controllers"
	models "lets/models"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/facebookgo/inject"
	gomock "github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/stretchr/testify/assert"
)

// Mock of ProductRepo interface
type MockProductRepo struct {
	ctrl     *gomock.Controller
	recorder *_MockProductRepoRecorder
}

// Recorder for MockProductRepo (not exported)
type _MockProductRepoRecorder struct {
	mock *MockProductRepo
}

func NewMockProductRepo(ctrl *gomock.Controller) *MockProductRepo {
	mock := &MockProductRepo{ctrl: ctrl}
	mock.recorder = &_MockProductRepoRecorder{mock}
	return mock
}

func (_m *MockProductRepo) EXPECT() *_MockProductRepoRecorder {
	return _m.recorder
}

func (_m *MockProductRepo) FindAll() []models.Product {
	ret := _m.ctrl.Call(_m, "FindAll")
	ret0, _ := ret[0].([]models.Product)
	return ret0
}

func (_mr *_MockProductRepoRecorder) FindAll() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FindAll")
}

func (_m *MockProductRepo) Find(id int) (*models.Product, error) {
	ret := _m.ctrl.Call(_m, "Find", id)
	ret0, _ := ret[0].(*models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockProductRepoRecorder) Find(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Find", arg0)
}

func (_m *MockProductRepo) Save(_param0 *models.Product) (*models.Product, []error) {
	ret := _m.ctrl.Call(_m, "Save", _param0)
	ret0, _ := ret[0].(*models.Product)
	ret1, _ := ret[1].([]error)
	return ret0, ret1
}

func (_mr *_MockProductRepoRecorder) Save(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Save", arg0)
}

func TestIndexMock(t *testing.T) {
	gctrl := gomock.NewController(t)
	defer gctrl.Finish()
	mockRepo := NewMockProductRepo(gctrl)
	mockRepo.EXPECT().FindAll().Return(
		[]models.Product{
			models.Product{
				Code:  "ABC",
				Price: 10,
			},
			models.Product{
				Code:  "DEF",
				Price: 20,
			},
		},
	)

	var ctrl controllers.ProductController
	inject.Populate(&ctrl, mockRepo)

	e := echo.New()
	req := new(http.Request)
	rec := httptest.NewRecorder()
	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
	c.SetPath("/")

	if assert.NoError(t, ctrl.Index(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var products []models.Product
		json.Unmarshal(([]byte)(rec.Body.String()), &products)
		assert.Equal(t, 2, len(products))
		assert.Equal(t, "ABC", products[0].Code)
		assert.Equal(t, 10, (int)(products[0].Price))

		assert.Equal(t, "DEF", products[1].Code)
		assert.Equal(t, 20, (int)(products[1].Price))
	}
}

func TestPostMock(t *testing.T) {
	gctrl := gomock.NewController(t)
	defer gctrl.Finish()
	mockRepo := NewMockProductRepo(gctrl)

	mockRepo.EXPECT().Save(gomock.Any()).Do(func(m *models.Product) {
		assert.Equal(t, "B1", m.Code)
		assert.Equal(t, 200, m.Price)
	}).Return(
		&models.Product{
			Id:    1,
			Code:  "B1",
			Price: 200,
		}, nil,
	)
	var ctrl controllers.ProductController
	inject.Populate(&ctrl, mockRepo)

	f := make(url.Values)
	f.Set("Code", "B1")
	f.Set("Price", "200")

	e := echo.New()
	req, _ := http.NewRequest(echo.POST, "/", strings.NewReader(f.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()
	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))

	if assert.NoError(t, ctrl.Post(c)) {
		var product models.Product
		json.Unmarshal(([]byte)(rec.Body.String()), &product)
		assert.Equal(t, "B1", product.Code)
		assert.Equal(t, 200, product.Price)
		assert.Equal(t, 1, product.Id)
	} else {
		t.Error("Should not return error")
	}
}
