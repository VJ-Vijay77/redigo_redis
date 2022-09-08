package routes

import (
	"github.com/labstack/echo/v4"
)

func EchoService() *echo.Echo {

	e := echo.New()
	return e
}
