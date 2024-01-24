package main

import (
	"github.com/aaron-smits/templ-starter/handlers"
	// "github.com/aaron-smits/templ-starter/db"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// supa "github.com/nedpals/supabase-go"
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
	app.Pre(middleware.RemoveTrailingSlash())
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
	todoHandler := handlers.TodoHandler{}
	// Groups
	auth := app.Group("/api/auth")
	todo := app.Group("/api/todo")
	// Routes

	// app.Use(withUser)
	app.GET("/", homeHandler.HandleHomeShow)

	auth.POST("/api/auth/login/github", userHandler.HandleUserLoginPost)
	auth.GET("/api/auth/login/callback", userHandler.HandleUserLoginCallback)
	auth.POST("/api/auth/logout", userHandler.HandleUserLogoutPost)

	todo.GET("/api/todo", todoHandler.HandleTodoGet)
	todo.POST("/api/todo", todoHandler.HandleTodoPost)
	todo.PUT("/api/todo/:id", todoHandler.HandleTodoPut)
	todo.DELETE("/api/todo/:id", todoHandler.HandleTodoDelete)

	app.Logger.Fatal(app.Start(":5173"))
}

// func withUser(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		ctx := context.WithValue(c.Request().Context(), "user", 1)
// 		c.SetRequest(c.Request().WithContext(ctx))
// 		return next(c)
// 	}
// }
