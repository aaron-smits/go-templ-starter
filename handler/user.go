package handler

import (
	"net/http"
	"os"

	"github.com/aaron-smits/templ-starter/model"
	"github.com/aaron-smits/templ-starter/view/loginview"
	"github.com/aaron-smits/templ-starter/view/userview"
	"github.com/labstack/echo/v4"

	supa "github.com/nedpals/supabase-go"
)

type UserHandler struct {
}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	u := model.User{
		Email: "test@example.com",
		ID: 1,
	}
	return render(c, userview.Show(u))
}

func (h UserHandler) HandleUserLoginShow(c echo.Context) error {
	return render(c, loginview.Login())
}

func (h UserHandler) HandleUserLoginPost(c echo.Context) error {

	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	supabase := supa.CreateClient(supabaseURL, supabaseKey)

	ProviderSignInOptions := supa.ProviderSignInOptions{
		Provider: "github",
	}

	ProviderSignInDetails, err := supabase.Auth.SignInWithProvider(ProviderSignInOptions)
	if err != nil {
		return err
	}
	return c.Redirect(http.StatusFound, ProviderSignInDetails.URL)
}