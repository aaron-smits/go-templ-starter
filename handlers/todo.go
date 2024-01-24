package handlers

import (
	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
}

func (h TodoHandler) HandleTodoGet(c echo.Context) error {
	return nil
}

func (h TodoHandler) HandleTodoPost(c echo.Context) error {
	return nil

}

func (h TodoHandler) HandleTodoPut(c echo.Context) error {
	return nil

}

func (h TodoHandler) HandleTodoDelete(c echo.Context) error {
	return nil
}
