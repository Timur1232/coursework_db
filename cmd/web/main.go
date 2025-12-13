package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Timur1232/coursework_db/internal/db"
	"github.com/Timur1232/coursework_db/internal/handlers"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var DB *pgx.Conn

const (
	user     = "app"
	password = "app"
	host     = "127.0.0.1"
	port     = 5432
	database = "coursework"
)

func getConnString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, host, port, database)
}

func main() {
	DB, err := pgx.Connect(context.Background(), getConnString())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer DB.Close(context.Background())

	e := echo.New()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &db.DBContext{Context: c, DB: DB}
			return next(cc)
		}
	})
	e.Use(middleware.Logger())

	e.Static("/static", "static")

	e.GET("/admin", func(c echo.Context) error {
		return c.Redirect(http.StatusTemporaryRedirect, "/admin/vgk")
	})

	e.GET("/admin/:tableName", func(c echo.Context) error {
		return handlers.AdminPanel(c)
	})

	e.GET("/api/admin/:tableName/:page", func(c echo.Context) error {
		isHxReq := c.Request().Header.Get("HX-Request") == "true"
		page := c.Param("page")
		if page == "" || !isHxReq {
			tableName := c.Param("tableName")
			return c.Redirect(http.StatusFound, "/admin/" + tableName)
		}
		return handlers.AdminPanelPage(c)
	})

	e.GET("/login", func(c echo.Context) error {
		return handlers.Login(c)
	})

	e.POST("/logout", func(c echo.Context) error {
		// TODO:
		return handlers.Logout(c)
	})

	e.GET("/register", func(c echo.Context) error {
		return handlers.Register(c)
	})

	e.GET("/api/home/accidents", func(c echo.Context) error {
		return c.NoContent(http.StatusNoContent)
	})

	e.GET("/api/home/objects", func(c echo.Context) error {
		return c.NoContent(http.StatusNoContent)
	})

	e.Logger.Fatal(e.Start(":42069"))

}
