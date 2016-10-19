package controllers

import (
	"lets/repo"
	"net/http"

	"github.com/labstack/echo"
)

type ProductController struct {
	ProductRepo repo.ProductRepo `inject:""`
}

func (ctrl *ProductController) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, ctrl.ProductRepo.FindAll())
}

func (ctrl *ProductController) IndexHtml(c echo.Context) error {
	params := map[string]interface{}{}
	params["products"] = ctrl.ProductRepo.FindAll()
	return c.Render(http.StatusOK, "product/index.html", params)
}
