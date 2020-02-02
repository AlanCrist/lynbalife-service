package api

import (
	"../api/handlers"

	"github.com/labstack/echo"
)

func CookieGroup(c *echo.Group) {
	c.GET("/main", handlers.MainCookie)
}
