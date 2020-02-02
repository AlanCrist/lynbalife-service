package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetJwtMiddlewares(j *echo.Group) {
	j.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte("mySecret"),
		TokenLookup:   "cookie:JWTCookie",
	}))
}
