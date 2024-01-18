package main

import (
	"context"
	"github.com/aaron-smits/templ-starter/handler"
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
	app.Use(withUser)
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		app.Logger.Fatal(err)
	}
	
	// DB
	// db, err := db.NewPostgresDB()
	// if err != nil {
	// 	app.Logger.Fatal(err)
	// }
	// app.Logger.Fatal(db.CreateUserTable())

	userHandler := handler.UserHandler{}
	homeHandler := handler.HomeHandler{}
	// Routes
	app.GET("/", homeHandler.HandleHomeShow)
	app.GET("/users/:id", userHandler.HandleUserShow)
	app.GET("/login", userHandler.HandleUserLoginShow)
	app.POST("/login", userHandler.HandleUserLoginPost)
	app.Logger.Fatal(app.Start(":5173"))
}


func withUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), "user", 1)
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}