package main

import (
	"fmt"
	"lets/controllers"
	"lets/db"
	"lets/models"
	"lets/repo"
	"os"

	"github.com/facebookgo/inject"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

func main() {
	e := echo.New()
	DB := db.InitDB()
	defer DB.Close()

	DB.Create(&models.Product{Code: "L1212", Price: 1000})

	var g inject.Graph

	var productCtrl controllers.ProductController
	var productRepo repo.ProductRepoImpl

	err := g.Provide(
		&inject.Object{Value: &productCtrl},
		&inject.Object{Value: &productRepo},
		&inject.Object{Value: DB},
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := g.Populate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	e.GET("/", productCtrl.Index)

	e.Run(standard.New(":1323"))
}
