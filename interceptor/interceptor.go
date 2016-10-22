package interceptor

import (
	"fmt"

	"github.com/labstack/echo"
)

func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println(c.Path())
		// return echo.ErrUnauthorized
		return next(c)
	}
}
