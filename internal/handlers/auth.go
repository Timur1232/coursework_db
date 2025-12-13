package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Timur1232/coursework_db/internal/db"
	"github.com/Timur1232/coursework_db/views"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	loginForm := views.LoginForm()
	if c.Request().Header.Get("HX-Request") == "true" {
		return RenderPage(c, loginForm)
	}
	page := views.Layout("Горноспасательная служба", loginForm, c.(*db.DBContext).User)
	return RenderPage(c, page)
}

func Register(c echo.Context) error {
	registerForm := views.RegisterForm()
	if c.Request().Header.Get("HX-Request") == "true" {
		return RenderPage(c, registerForm)
	}
	page := views.Layout("Горноспасательная служба", registerForm, c.(*db.DBContext).User)
	return RenderPage(c, page)
}

func PostLogin(c echo.Context) error {
	login := c.FormValue("username")
	password := c.FormValue("password")

	user, err := db.FindUserByLogin(c.(*db.DBContext).DB, login)
	if err != nil || user.Password != password {
		return c.HTML(http.StatusOK, `<div class="message message-error">Неверный логин или пароль</div>`)
	}

	cookie := new(http.Cookie)
	cookie.Name = "user_id"
	cookie.Value = strconv.FormatUint(user.IdUser, 10)
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.MaxAge = 86400 // 24 часа
	c.SetCookie(cookie)

	c.Response().Header().Set("HX-Redirect", "/")
	return c.NoContent(http.StatusOK)
}

func PostRegister(c echo.Context) error {
	login := c.FormValue("username")
	password := c.FormValue("password")

	existingUser, _ := db.FindUserByLogin(c.(*db.DBContext).DB, login)
	if existingUser != nil {
		return c.HTML(http.StatusOK, `<div class="message message-error">Пользователь с таким логином уже существует</div>`)
	}

	var newUserID uint64
	err := c.(*db.DBContext).DB.QueryRow(context.Background(),
		"INSERT INTO users (login, password, role) VALUES ($1, $2, $3) RETURNING id_user",
		login, password, db.Role_Guest).Scan(&newUserID)
	if err != nil {
		return c.HTML(http.StatusOK, `<div class="message message-error">Внутренняя ошибка сервера</div>`)
	}

	cookie := new(http.Cookie)
	cookie.Name = "user_id"
	cookie.Value = strconv.FormatUint(newUserID, 10)
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.MaxAge = 86400
	c.SetCookie(cookie)

	c.Response().Header().Set("HX-Redirect", "/")
	return c.NoContent(http.StatusOK)
}

func PostLogout(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "user_id"
	cookie.Value = ""
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.MaxAge = -1
	c.SetCookie(cookie)

	c.Response().Header().Set("HX-Redirect", "/login")
	return c.NoContent(http.StatusOK)
}
