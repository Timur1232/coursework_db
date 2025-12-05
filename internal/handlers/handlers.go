package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/a-h/templ"

	"github.com/Timur1232/coursework_db/views"
)

func RenderPage(c echo.Context, component templ.Component) error {
	page := views.Page(component)
	return page.Render(c.Request().Context(), c.Response().Writer)
}
