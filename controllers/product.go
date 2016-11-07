package controllers

import (
	"lets/models"
	"lets/repo"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type ProductController struct {
	ProductRepo repo.ProductRepo `inject:""`
}

func (ctrl *ProductController) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, ctrl.ProductRepo.FindAll())
}

func (ctrl *ProductController) Post(c echo.Context) error {
	price, _ := strconv.Atoi(c.FormValue("Price"))
	m, _ := ctrl.ProductRepo.Save(&models.Product{
		Code:  c.FormValue("Code"),
		Price: price,
	})
	return c.JSON(http.StatusCreated, m)
}

func (ctrl *ProductController) IndexHtml(c echo.Context) error {
	params := map[string]interface{}{}
	params["products"] = ctrl.ProductRepo.FindAll()
	return c.Render(http.StatusOK, "product/index.html", params)
}
