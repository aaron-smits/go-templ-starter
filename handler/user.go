package handler

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"

	supa "github.com/nedpals/supabase-go"
)

type UserHandler struct {
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

	c.SetCookie(&http.Cookie{
		Name:     "code_verifier",
		Value:    ProviderSignInDetails.CodeVerifier,
		HttpOnly: true,
		Domain:   "localhost",
		Path:     "/",
	})
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
	c.SetCookie(&http.Cookie{
		Name:    "code_verifier",
		Value:   "",
		Expires: time.Unix(0, 0),
		Secure: true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
	c.SetCookie(&http.Cookie{
		Name:     "access_token",
		Value:    ExchangeCodeDetails.AccessToken,
		Domain:   "localhost",
		Secure: true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
	c.SetCookie(&http.Cookie{
		Name:   "refresh_token",
		Value:  ExchangeCodeDetails.RefreshToken,
		Domain: "localhost",
		Secure: true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,

	})
	c.SetCookie(&http.Cookie{
		Name: "user_id",
		Value: ExchangeCodeDetails.User.ID,
		Domain: "localhost",
		Secure: true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
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
		Expires: time.Unix(0, 0),
		Secure: true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
	c.SetCookie(&http.Cookie{
		Name:    "refresh_token",
		Value:   "",
		Expires: time.Unix(0, 0),
		Secure: true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	return c.Redirect(http.StatusFound, "/")
}
