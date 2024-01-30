package handlers

import (
	"github.com/aaron-smits/templ-starter/db"
	"github.com/aaron-smits/templ-starter/model"
	"github.com/aaron-smits/templ-starter/view/components"
	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	DB db.DB
}

func (h TodoHandler) HandleTodoPost(c echo.Context) error {
	// user := c.Get("user").(*model.User)
	userIdCookie, err := c.Cookie("user_id")
	if err != nil {
		return err
	}
	userId := userIdCookie.Value
	title := c.FormValue("title")
	body := c.FormValue("body")
	todo := model.Todo{
		UserID: userId,
		Title:  title,
		Body:   body,
		Done:   false,
	}

	// append the new todo to the list
	todoErr := h.DB.AddTodo(todo)
	if todoErr != nil {
		return todoErr
	}
	todoList, err := h.DB.GetTodoList()
	if err != nil {
		return err
	}
	return Render(c, components.TodoList(todoList))
}

func (h TodoHandler) HandleTodoDelete(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.String(400, "ID is required")
	}
	err := h.DB.DeleteTodo(id)
	if err != nil {
		return err
	}
	todoList, err := h.DB.GetTodoList()
	if err != nil {
		return err
	}
	return Render(c, components.TodoList(todoList))
}
