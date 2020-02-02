package api

import (
	"../api/handlers"
	"github.com/labstack/echo"
)

func MainGroup(e *echo.Echo) {
	e.GET("/login", handlers.Login)
	e.GET("/yallo", handlers.Yallo)
}
