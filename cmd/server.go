package main

import (
	"fmt"
	"os"

	"github.com/aaron-smits/templ-starter/db"
	"github.com/labstack/echo/v4"
)

type Server struct {
	app    *echo.Echo
	Config Config
	DB     db.DB
}

func NewServer() (*Server, error) {
	db, err := db.NewPostgresDB()
	if err != nil {
		fmt.Printf("Error connecting to database: %v", err)
		os.Exit(1)
	}
	// Create the todo table if it doesn't exist
	err = db.CreateTodoTable()
	if err != nil {
		fmt.Printf("Error creating todo table: %v", err)
		os.Exit(1)
	}
	return &Server{
		app:    echo.New(),
		Config: NewConfig(),
		DB:     db,
	}, nil
}

func NewServerWithConfig(config Config) (*Server, error) {
	return &Server{
		app:    echo.New(),
		Config: config,
	}, nil
}
