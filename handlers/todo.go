package handlers

import (
	"github.com/aaron-smits/templ-starter/model"
	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
}

var TodoList []*model.Todo = []*model.Todo{}

func (h TodoHandler) HandleTodosGet(c echo.Context) error {
	todo := model.Todo{
		ID:     "1",
		UserID: 1,
		Title:  "Test",
		Body:   "Test",
		Done:   false,
	}
	TodoList = append(TodoList, &todo)

	return c.JSON(200, TodoList)
}

func (h TodoHandler) HandleTodoPost(c echo.Context) error {
	todo := new(model.Todo)
	err := c.Bind(todo)
	if err != nil {
		return err
	}
	TodoList = append(TodoList, todo)
	return c.JSON(200, TodoList)
}

func (h TodoHandler) HandleTodoPut(c echo.Context) error {
	todo := new(model.Todo)
	err := c.Bind(todo)
	if err != nil {
		return err
	}
	// Find the todo in the list with the matching ID
	for i, t := range TodoList {
		if t.ID == todo.ID {
			TodoList[i] = todo
		}
	}
	return c.JSON(200, TodoList)
}

func (h TodoHandler) HandleTodoDelete(c echo.Context) error {
	// Iterate through the list and remove the todo with the matching ID
	for i, t := range TodoList {
		if t.ID == c.Param("id") {
			TodoList = append(TodoList[:i], TodoList[i+1:]...)
		}
	}

	return c.JSON(200, TodoList)
}
