package handlers

import (
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

// Use this to clear a cookie
func deleteCookie(c echo.Context, name string) {
	c.SetCookie(&http.Cookie{
		Name:     name,
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Domain:   "localhost",
		Path:     "/",
		Secure:   true,
	})
}

// Use this to set a cookie with a secure flag, http only, and a domain
func setCookie(c echo.Context, name string, value string) {
	c.SetCookie(&http.Cookie{
		Name:     name,
		Value:    value,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Domain:   "localhost",
		Path:     "/",
		Secure:   true,
	})
}
