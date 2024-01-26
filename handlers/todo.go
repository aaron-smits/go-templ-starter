package handlers

import (
	"github.com/aaron-smits/templ-starter/db"
	"github.com/aaron-smits/templ-starter/model"
	"github.com/aaron-smits/templ-starter/view/homeview"
	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
}

func (h TodoHandler) HandleTodoPost(c echo.Context) error {
	user := c.Get("user").(*model.User)
	userId := user.User.ID
	if userId == "" {
		return Render(c, homeview.LoggedOutHome())
	}
	title := c.FormValue("title")
	if title == "" {
		return Render(c, homeview.Home(user, db.TodoList))
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
	return Render(c, homeview.Home(user, db.TodoList))
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

// func (h TodoHandler) HandleTodoDelete(c echo.Context) error {
// 	// Iterate through the list and remove the todo with the matching ID
// 	for i, t := range db.TodoList {
// 		if t.ID == c.Param("id") {
// 			db.TodoList = append(db.TodoList[:i], db.TodoList[i+1:]...)
// 		}
// 	}

// 	return c.JSON(200, db.TodoList)
// }
