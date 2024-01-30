package main

import (
	"github.com/aaron-smits/templ-starter/handlers"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	s, err := NewServer()
	if err != nil {
		panic(err)
	}
	app := s.app
	// Middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
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
	todo.DELETE("/:id", handlers.WithAuth(todoHandler.HandleTodoDelete))

	app.Logger.Fatal(app.Start(":" + s.config.Port))
}
