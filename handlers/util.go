package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/a-h/templ"
	"github.com/aaron-smits/templ-starter/model"
	"github.com/aaron-smits/templ-starter/view/pages"
	"github.com/labstack/echo/v4"

	supa "github.com/nedpals/supabase-go"
)

func Render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func WithAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		supabaseURL := os.Getenv("SUPABASE_URL")
		supabaseKey := os.Getenv("SUPABASE_KEY")
		supabase := supa.CreateClient(supabaseURL, supabaseKey)

		token, err := c.Cookie("access_token")
		if err != nil || token == nil {
			return Render(c, pages.LoggedOutHome())
		}

		user, err := supabase.Auth.User(c.Request().Context(), token.Value)
		if err != nil {
			return Render(c, pages.LoggedOutHome())
		}
		if user == nil {
			return Render(c, pages.LoggedOutHome())
		}
		modelUser := &model.User{
			User: user,
		}

		c.Set("user", modelUser)
		return next(c)
	}
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
