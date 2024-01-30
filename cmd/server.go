package main

import (
	"github.com/labstack/echo/v4"
)

type Server struct {
	app    *echo.Echo
	config Config
}

func NewServer() (*Server, error) {
	return &Server{
		app:    echo.New(),
		config: NewConfig(),
	}, nil
}

func NewServerWithConfig(config Config) (*Server, error) {
	return &Server{
		app:    echo.New(),
		config: config,
	}, nil
}
