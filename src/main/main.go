package main

import (
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	fmt.Println("welcome to the server")

	e := echo.New()

	e.Use(ServerHeader)

	adminGroup := e.Group("/admin")
	cookieGroup := e.Group("/cookie")

	adminGroup.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}]  ${status}  ${method}  ${host}${path}  ${latency_human}` + "\n",
	}))

	adminGroup.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte("joe")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("secret")) == 1 {
			return true, nil
		}
		return false, nil
	}))

	cookieGroup.Use(checkCookie)

	cookieGroup.GET("/main", mainCookie)
	adminGroup.GET("/main", mainAdmin)

	e.GET("/login", login)
	e.GET("/", yello)

	e.Start(":8000")
}

func yello(c echo.Context) error {
	return c.String(http.StatusOK, "yello from the web side")
}

func mainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "horay you are on the secret admin main page")
}

func mainCookie(c echo.Context) error {
	return c.String(http.StatusOK, "horay you are on the secret cookie main page")
}

func login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	if username == "joe" && password == "secret" {
		cookie := &http.Cookie{}

		cookie.Name = "sessionID"
		cookie.Value = "some_string"
		cookie.Expires = time.Now().Add(48 * time.Hour)

		c.SetCookie(cookie)

		return c.String(http.StatusOK, "You were logged in!")
	}

	return c.String(http.StatusUnauthorized, "Your username or password were wrong")
}

func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("sessionID")
		if err != nil {

			if strings.Contains(err.Error(), "named cookie not present") {
				return c.String(http.StatusUnauthorized, "you dont have any cookie")
			}

			log.Println(err)
			return err
		}

		if cookie.Value == "some_string" {
			return next(c)
		}

		return c.String(http.StatusUnauthorized, "you dont have the right cookie, cookie")
	}
}

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "LynbaLife/1.0")
		c.Response().Header().Set("notReallyHeader", "thisHaveNoMeaning")

		return next(c)
	}
}
