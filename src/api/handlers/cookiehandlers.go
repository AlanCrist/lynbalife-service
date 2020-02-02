package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func MainCookie(c echo.Context) error {
	return c.String(http.StatusOK, "horay you are on the secret cookie main page")
}
