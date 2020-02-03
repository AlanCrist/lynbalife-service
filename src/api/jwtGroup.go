package api

import (
	"api/handlers"

	"github.com/labstack/echo"
)

func JwtGroup(j *echo.Group) {
	j.GET("/main", handlers.MainJwt)

}
