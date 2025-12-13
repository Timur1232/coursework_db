package handlers

import (
	"github.com/Timur1232/coursework_db/internal/db"
	"github.com/Timur1232/coursework_db/views"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	loginForm := views.LoginForm()
	if c.Request().Header.Get("HX-Request") == "true" {
		return RenderPage(c, loginForm)
	}
	page := views.Layout("Горноспасательная служба", loginForm, &db.Users{IdUser: 69, Login: "Timur Baimuradov", Password: "a$$word", Role: db.Role_Admin})
	return RenderPage(c, page)
}

func Register(c echo.Context) error {
	registerForm := views.RegisterForm()
	if c.Request().Header.Get("HX-Request") == "true" {
		return RenderPage(c, registerForm)
	}
	page := views.Layout("Горноспасательная служба", registerForm, &db.Users{IdUser: 69, Login: "Timur Baimuradov", Password: "a$$word", Role: db.Role_Admin})
	return RenderPage(c, page)
}

func Logout(c echo.Context) error {
	logout := views.LogoutNotification()
	if c.Request().Header.Get("HX-Request") == "true" {
		return RenderPage(c, logout)
	}
	page := views.Layout("Горноспасательная служба", logout, &db.Users{IdUser: 69, Login: "Timur Baimuradov", Password: "a$$word", Role: db.Role_Admin})
	return RenderPage(c, page)
}
