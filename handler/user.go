package handler

import (
	"net/http"
	"os"
	"time"

	"github.com/aaron-smits/templ-starter/model"
	"github.com/aaron-smits/templ-starter/view/userview"
	"github.com/labstack/echo/v4"

	supa "github.com/nedpals/supabase-go"
)

type UserHandler struct {
}

type CodeVerifier struct {
	CodeVerifier string `json:"code_verifier"`
}

// Create storage for the code verifier to pass it between login and callback
var codeVerifier CodeVerifier

func (h UserHandler) HandleUserShow(c echo.Context) error {
	u := model.User{
		Email: "test@example.com",
		ID:    1,
	}
	return render(c, userview.Show(u))
}

func (h UserHandler) HandleUserLoginPost(c echo.Context) error {
	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	supabase := supa.CreateClient(supabaseURL, supabaseKey)

	ProviderSignInOptions := supa.ProviderSignInOptions{
		Provider:   "github",
		FlowType:   "pkce",
		RedirectTo: "http://localhost:5173/login/callback",
	}

	ProviderSignInDetails, err := supabase.Auth.SignInWithProvider(ProviderSignInOptions)
	if err != nil {
		return err
	}
	codeVerifier = CodeVerifier{
		CodeVerifier: ProviderSignInDetails.CodeVerifier,
	}
	return c.Redirect(http.StatusFound, ProviderSignInDetails.URL)
}

func (h UserHandler) HandleUserLoginCallback(c echo.Context) error {
	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	supabase := supa.CreateClient(supabaseURL, supabaseKey)
	ExchangeCodeOptions := supa.ExchangeCodeOpts{
		AuthCode:     c.QueryParam("code"),
		CodeVerifier: codeVerifier.CodeVerifier,
	}
	ExchangeCodeDetails, err := supabase.Auth.ExchangeCode(c.Request().Context(), ExchangeCodeOptions)
	if err != nil {
		return err
	}
	c.SetCookie(&http.Cookie{
		Name:     "access_token",
		Value:    ExchangeCodeDetails.AccessToken,
		HttpOnly: true,
		Domain:   "localhost",
		Path:     "/",
	})
	c.SetCookie(&http.Cookie{
		Name:  "refresh_token",
		Value: ExchangeCodeDetails.RefreshToken,
	})
	return c.Redirect(http.StatusFound, "/")
}

func (h UserHandler) HandleUserLogoutPost(c echo.Context) error {
	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	supabase := supa.CreateClient(supabaseURL, supabaseKey)
	token, err := c.Cookie("access_token")
	if err != nil {
		return err
	}
	err = supabase.Auth.SignOut(c.Request().Context(), token.Value)
	if err != nil {
		return err
	}
	c.SetCookie(&http.Cookie{
		Name:    "access_token",
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),
	})
	c.SetCookie(&http.Cookie{
		Name:    "refresh_token",
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),
	})

	return c.Redirect(http.StatusFound, "/")
}
