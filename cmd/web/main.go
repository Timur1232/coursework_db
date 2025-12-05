package main

import (
	// "context"
	// "fmt"

	// "github.com/jackc/pgx/v5"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// "github.com/Timur1232/coursework_db/internal/handlers"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Static("/", "static")

	e.GET("/admin", func(c echo.Context) error {
		return c.NoContent(501)
	})

	e.Logger.Fatal(e.Start(":42069"))
}
