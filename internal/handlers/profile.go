package handlers

import (
	"github.com/Timur1232/coursework_db/internal/db"
	"github.com/Timur1232/coursework_db/views"
	"github.com/labstack/echo/v4"
)

func Profile(c echo.Context) error {
	user := c.(*db.DBContext).User
	profilePage := views.ProfilePage(user)
	if c.Request().Header.Get("HX-Request") == "true" {
		return RenderPage(c, profilePage)
	}
	page := views.Layout("Личный кабинет", profilePage, user)
	return RenderPage(c, page)
}
