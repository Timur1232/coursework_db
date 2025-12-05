package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/a-h/templ"
)

func RenderPage(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response().Writer)
}
