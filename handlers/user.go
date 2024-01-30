package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	supa "github.com/nedpals/supabase-go"

	"github.com/aaron-smits/templ-starter/db"
)

type UserHandler struct {
	DB db.DB
}

func (h UserHandler) HandleUserLoginPost(c echo.Context) error {
	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	baseUrl := os.Getenv("BASE_URL")
	port := os.Getenv("PORT")
	if baseUrl == "http://localhost" {
		baseUrl = fmt.Sprintf("%s:%s", baseUrl, port)
	}
	supabase := supa.CreateClient(supabaseURL, supabaseKey)

	ProviderSignInOptions := supa.ProviderSignInOptions{
		Provider:   "github",
		FlowType:   "pkce",
		RedirectTo: baseUrl + "/api/auth/login/callback",
	}

	ProviderSignInDetails, err := supabase.Auth.SignInWithProvider(ProviderSignInOptions)
	if err != nil {
		return err
	}
	setCookie(c, "code_verifier", ProviderSignInDetails.CodeVerifier)
	return c.Redirect(http.StatusFound, ProviderSignInDetails.URL)
}

func (h UserHandler) HandleUserLoginCallback(c echo.Context) error {
	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	supabase := supa.CreateClient(supabaseURL, supabaseKey)
	codeVerifier, err := c.Cookie("code_verifier")
	if err != nil {
		return err
	}
	if codeVerifier == nil {
		// return a new error
		return fmt.Errorf("code verifier cookie not found")
	}

	ExchangeCodeOptions := supa.ExchangeCodeOpts{
		AuthCode:     c.QueryParam("code"),
		CodeVerifier: codeVerifier.Value,
	}
	ExchangeCodeDetails, err := supabase.Auth.ExchangeCode(c.Request().Context(), ExchangeCodeOptions)
	if err != nil {
		return err
	}
	deleteCookie(c, "code_verifier")
	setCookie(c, "access_token", ExchangeCodeDetails.AccessToken)
	setCookie(c, "refresh_token", ExchangeCodeDetails.RefreshToken)
	setCookie(c, "user_id", ExchangeCodeDetails.User.ID)
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
	deleteCookie(c, "access_token")

	return c.Redirect(http.StatusFound, "/")
}
