package api

import (
	"github.com/VJ-Vijay77/redis_redigo_pkg/controllers"
	"github.com/labstack/echo/v4"
)



func EchoApi(e *echo.Echo) {
	e.GET("/home",controllers.Home)
	e.POST("/add",controllers.AddUser)
	e.GET("/getpass/:name",controllers.GetPass)
}

