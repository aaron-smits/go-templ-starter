package handler

import (
	"github.com/aaron-smits/templ-starter/view/userview"
	"github.com/aaron-smits/templ-starter/model"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	u := model.User{
		Email: "test@example.com",
		ID: 1,
	}
	return render(c, userview.Show(u))
}