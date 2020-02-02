package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func Yallo(c echo.Context) error {
	return c.String(http.StatusOK, "yallo from the web side")
}
