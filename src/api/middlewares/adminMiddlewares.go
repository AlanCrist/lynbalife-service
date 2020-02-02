package middlewares

import (
	"crypto/subtle"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetAdminMiddlewares(g *echo.Group) {
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}]  ${status}  ${method}  ${host}${path}  ${latency_human}` + "\n",
	}))

	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte("joe")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("secret")) == 1 {
			return true, nil
		}
		return false, nil
	}))
}
