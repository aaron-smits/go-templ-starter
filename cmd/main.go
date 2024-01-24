package main

import (
	"os"

	"github.com/aaron-smits/templ-starter/handlers"
	// "github.com/aaron-smits/templ-starter/db"
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
	app.Use(withUser)
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		app.Logger.Fatal(err)
	}
	// app.Pre(middleware.RemoveTrailingSlash())
	// app.Pre(middleware.HTTPSRedirect())
	// app.Pre(middleware.HTTPSNonWWWRedirect())
	// app.Pre(middleware.NonWWWRedirect())
	// app.Use(middleware.Secure())
	// DB
	// db, err := db.NewPostgresDB()
	// if err != nil {
	// 	app.Logger.Fatal(err)
	// }
	// app.Logger.Fatal(db.CreateUserTable())

	userHandler := handlers.UserHandler{}
	homeHandler := handlers.HomeHandler{}
	// todoHandler := handlers.TodoHandler{}
	// Groups
	auth := app.Group("/api/auth")
	// todo := app.Group("/api/todo")
	// Routes

	// app.Use(withUser)
	app.GET("/", homeHandler.HandleHomeShow)

	auth.POST("/login/github", userHandler.HandleUserLoginPost)
	auth.GET("/login/callback", userHandler.HandleUserLoginCallback)
	auth.POST("/logout", userHandler.HandleUserLogoutPost)

	// todo.GET("/", todoHandler.HandleTodoGet)
	// todo.POST("/", todoHandler.HandleTodoPost)
	// todo.PUT("/:id", todoHandler.HandleTodoPut)
	// todo.DELETE("/:id", todoHandler.HandleTodoDelete)

	app.Logger.Fatal(app.Start(":5173"))
}

func withUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		supabaseURL := os.Getenv("SUPABASE_URL")
		supabaseKey := os.Getenv("SUPABASE_KEY")
		supabase := supa.CreateClient(supabaseURL, supabaseKey)
		accessToken, err := c.Cookie("access_token")
		if err != nil {
			return nil
		}
		if accessToken == nil {
			return nil
		}
		user, err := supabase.Auth.User(c.Request().Context(), accessToken.Value)
		if err != nil {
			return nil
		}
		c.Set("user", user)
		return next(c)
	}
}
