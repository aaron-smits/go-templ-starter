package handler

import (
	"github.com/aaron-smits/templ-starter/view/homeview"
	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
}

func (h HomeHandler) HandleHomeShow(c echo.Context) error {
	return render(c, homeview.Home())
}