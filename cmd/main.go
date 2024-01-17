package main

import (
	"context"
	"fmt"

	"github.com/aaron-smits/templ-starter/handler"
	"github.com/labstack/echo/v4"
)

type DB struct {
}

func main() {
	app := echo.New()
	userHandler := handler.UserHandler{}
	app.Use(withUser)
	app.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	app.GET("/users/:id", userHandler.HandleUserShow)

	app.Start(":5173")
	fmt.Println("Hello, World!")
}


func withUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), "user", 1)
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}