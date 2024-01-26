package main

import (
	"github.com/aaron-smits/templ-starter/handlers"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	app.GET("/", handlers.WithAuth(homeHandler.HandleHomeShow))

	auth.POST("/login/github", userHandler.HandleUserLoginPost)
	auth.GET("/login/callback", userHandler.HandleUserLoginCallback)
	auth.POST("/logout", handlers.WithAuth(userHandler.HandleUserLogoutPost))

	todo.POST("/", handlers.WithAuth(todoHandler.HandleTodoPost))
	todo.PUT("/:id", handlers.WithAuth(todoHandler.HandleTodoPut))
	// todo.DELETE("/:id", todoHandler.HandleTodoDelete)

	app.Logger.Fatal(app.Start(":5173"))
}
