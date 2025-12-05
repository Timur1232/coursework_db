package main

import (
	"context"
	"fmt"

	"github.com/Timur1232/coursework_db/internal/db"
	"github.com/Timur1232/coursework_db/views"

	"github.com/a-h/templ"
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

func RenderPage(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func main() {
	DB, err := pgx.Connect(context.Background(), getConnString())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer DB.Close(context.Background())

	e := echo.New()
	e.Use(middleware.Logger())

	e.Static("/", "static")

	e.GET("/admin", func(c echo.Context) error {
		return c.NoContent(501)
	})

	e.GET("/admin/vgk", func(c echo.Context) error {
		res, err := db.Query[db.Vgk](DB, context.Background(), "SELECT * FROM vgk")
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		table := views.Table(views.TableListVgk(res))
		page := views.Page(table)
		return RenderPage(c, page)
	})

	e.Logger.Fatal(e.Start(":42069"))

}
