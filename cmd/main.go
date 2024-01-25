package main

import (
	"os"

	"github.com/aaron-smits/templ-starter/handlers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	supa "github.com/nedpals/supabase-go"
)

func main() {
	app := echo.New()
	// Middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		app.Logger.Fatal(err)
	}

	// Handlers
	userHandler := handlers.UserHandler{}
	homeHandler := handlers.HomeHandler{}
	todoHandler := handlers.TodoHandler{}
	// Groups
	auth := app.Group("/api/auth")
	todo := app.Group("/api/todo")

	// Routes

	app.GET("/", homeHandler.HandleHomeShow)

	auth.POST("/login/github", userHandler.HandleUserLoginPost)
	auth.GET("/login/callback", userHandler.HandleUserLoginCallback)
	auth.POST("/logout", withAuth(userHandler.HandleUserLogoutPost))

	todo.GET("/all", withAuth(todoHandler.HandleTodosGet))
	todo.POST("/", withAuth(todoHandler.HandleTodoPost))
	todo.PUT("/:id", withAuth(todoHandler.HandleTodoPut))
	// todo.DELETE("/:id", todoHandler.HandleTodoDelete)

	app.Logger.Fatal(app.Start(":5173"))
}

// func withUser(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		supabaseURL := os.Getenv("SUPABASE_URL")
// 		supabaseKey := os.Getenv("SUPABASE_KEY")
// 		supabase := supa.CreateClient(supabaseURL, supabaseKey)
// 		accessToken, err := c.Cookie("access_token")
// 		if err != nil {
// 			return err
// 		}
// 		if accessToken == nil {
// 			return err
// 		}
// 		user, err := supabase.Auth.User(c.Request().Context(), accessToken.Value)
// 		if err != nil {
// 			return err
// 		}
// 		c.Set("user", user)
// 		return next(c)
// 	}
// }

func withAuth(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authCookie, err := c.Cookie("access_token")
		homeHandler := handlers.HomeHandler{}

		if err != nil {
			return c.String(500, "user is not set")
		}
		if authCookie == nil {
			return homeHandler.HandleHomeShow(c)
		}
		supaURL := os.Getenv("SUPABASE_URL")
		supaKey := os.Getenv("SUPABASE_KEY")
		supa := supa.CreateClient(supaURL, supaKey)

		user, err := supa.Auth.User(c.Request().Context(), authCookie.Value)
		if err != nil {
			return c.String(401, "Unauthorized")
		}
		if user == nil {
			return homeHandler.HandleHomeShow(c)
		}
		return handlerFunc(c)
	}
}
