package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aaron-smits/templ-starter/handler"
	"github.com/aaron-smits/templ-starter/model"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type DB struct {
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := model.NewPostgresDB()
	if err != nil {
		log.Fatalf("Error creating database: %v", err)
	}
	if err = db.CreateUserTable(); err != nil {
		log.Fatalf("Error creating user table: %v", err)
	}
	fmt.Println("Database created")
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