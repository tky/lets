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
