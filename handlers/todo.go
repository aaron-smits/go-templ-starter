package handlers

import (
	"os"

	"github.com/aaron-smits/templ-starter/db"
	"github.com/aaron-smits/templ-starter/model"
	"github.com/aaron-smits/templ-starter/view/homeview"
	"github.com/labstack/echo/v4"

	supa "github.com/nedpals/supabase-go"
)

type TodoHandler struct {
}

func (h TodoHandler) HandleTodosGet(c echo.Context) error {
	return c.JSON(200, db.TodoList)
}

func (h TodoHandler) HandleTodoPost(c echo.Context) error {
	userId, err := c.Cookie("user_id")
	if err != nil {
		return err
	}
	title := c.FormValue("title")
	body := c.FormValue("body")
	todoUserId := userId.Value
	id := len(db.TodoList) + 1
	todo := model.Todo{
		ID:     id,
		UserID: todoUserId,
		Title:  title,
		Body:   body,
		Done:   false,
	}

	// append the new todo to the list
	db.TodoList = append(db.TodoList, todo)
	supaURL := os.Getenv("SUPABASE_URL")
	supaKey := os.Getenv("SUPABASE_KEY")
	supa := supa.CreateClient(supaURL, supaKey)
	token, err := c.Cookie("access_token")
	if err != nil {
		return err
	}
	user, err := supa.Auth.User(c.Request().Context(), token.Value)
	if err != nil {
		return err
	}
	return render(c, homeview.Home(user, db.TodoList))
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
