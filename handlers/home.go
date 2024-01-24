package handlers

import (
	"os"

	"github.com/aaron-smits/templ-starter/view/homeview"
	"github.com/labstack/echo/v4"

	supa "github.com/nedpals/supabase-go"
)

type HomeHandler struct {
}

func (h HomeHandler) HandleHomeShow(c echo.Context) error {

	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	supabase := supa.CreateClient(supabaseURL, supabaseKey)
	token, err := c.Cookie("access_token")
	if token == nil {
		return render(c, homeview.Home(""))
	}
	if err != nil {
		return render(c, homeview.Home(""))
	}
	user, err := supabase.Auth.User(c.Request().Context(), token.Value)

	if err != nil {
		return render(c, homeview.Home(""))
	}

	return render(c, homeview.Home(user.Email))
}
