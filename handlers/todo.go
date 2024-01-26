package handlers

import (
	"strconv"

	"github.com/aaron-smits/templ-starter/db"
	"github.com/aaron-smits/templ-starter/model"
	"github.com/aaron-smits/templ-starter/view/components"
	"github.com/aaron-smits/templ-starter/view/pages"
	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
}

func (h TodoHandler) HandleTodoPost(c echo.Context) error {
	user := c.Get("user").(*model.User)
	userId := user.User.ID
	title := c.FormValue("title")
	if title == "" {
		return Render(c, pages.Home(user, db.TodoList))
	}
	body := c.FormValue("body")
	id := len(db.TodoList) + 1
	todo := model.Todo{
		ID:     id,
		UserID: userId,
		Title:  title,
		Body:   body,
		Done:   false,
	}

	// append the new todo to the list
	db.TodoList = append(db.TodoList, todo)
	return Render(c, components.TodoList(db.TodoList))
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
	return Render(c, components.TodoList(db.TodoList))
}

func (h TodoHandler) HandleTodoDelete(c echo.Context) error {
	id := c.Param("id")
	// convert id to int
	strId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	// Find the todo in the list with the matching ID
	for i, t := range db.TodoList {
		if t.ID == strId {
			db.TodoList = append(db.TodoList[:i], db.TodoList[i+1:]...)
		}
	}
	return Render(c, components.TodoList(db.TodoList))
}
