package handlers

import (
	"github.com/aaron-smits/templ-starter/db"
	"github.com/aaron-smits/templ-starter/model"
	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
}

func (h TodoHandler) HandleTodosGet(c echo.Context) error {
	todo := model.Todo{
		ID:     "1",
		UserID: 1,
		Title:  "Test",
		Body:   "Test",
		Done:   false,
	}
	db.TodoList = append(db.TodoList, todo)

	return c.JSON(200, db.TodoList)
}

func (h TodoHandler) HandleTodoPost(c echo.Context) error {
	todo := new(model.Todo)
	err := c.Bind(todo)
	if err != nil {
		return err
	}
	db.TodoList = append(db.TodoList, *todo)
	return c.Redirect(302, "/")
}

func (h TodoHandler) HandleTodoPut(c echo.Context) error {
	todo := new(model.Todo)
	err := c.Bind(todo)
	if err != nil {
		return err
	}
	// Find the todo in the list with the matching ID
	for i, t := range db.TodoList {
		if t.ID == todo.ID {
			db.TodoList[i] = *todo
		}
	}
	return c.JSON(200, db.TodoList)
}

func (h TodoHandler) HandleTodoDelete(c echo.Context) error {
	// Iterate through the list and remove the todo with the matching ID
	for i, t := range db.TodoList {
		if t.ID == c.Param("id") {
			db.TodoList = append(db.TodoList[:i], db.TodoList[i+1:]...)
		}
	}

	return c.JSON(200, db.TodoList)
}
