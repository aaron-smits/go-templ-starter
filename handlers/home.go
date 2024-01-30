package handlers

import (
	"github.com/aaron-smits/templ-starter/db"
	"github.com/aaron-smits/templ-starter/model"
	"github.com/aaron-smits/templ-starter/view/pages"
	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
}

func (h HomeHandler) HandleHomeShow(c echo.Context) error {
	user := c.Get("user").(*model.User)
	return Render(c, pages.Home(user, db.TodoList))
}
