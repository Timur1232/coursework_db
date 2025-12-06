package handlers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Timur1232/coursework_db/internal/db"
	"github.com/Timur1232/coursework_db/views"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func RenderPage(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response().Writer)
}

const (
	pageSize = 10
)

func AdminPanel[T any](c echo.Context, tableName string, listComp func([]T) templ.Component, sortComp func() templ.Component) error {
	DB := c.(*db.DBContext).DB

	query := fmt.Sprintf("SELECT * FROM %s LIMIT $1", tableName)

	res, err := db.Query[T](DB, context.Background(), query, pageSize)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	query = fmt.Sprintf("SELECT count(*) FROM %s", tableName)

	var count int
	DB.QueryRow(context.Background(), query).Scan(&count)

	list := listComp(res)
	sort := sortComp()
	admin := views.AdminPanel(tableName, 1, pageSize, pageSize >= count, list, sort)
	page := views.Layout("Горноспасательная служба", admin)
	return RenderPage(c, page)
}

func AdminPanelPage[T any](c echo.Context, tableName string, listComp func([]T) templ.Component) error {
	DB := c.(*db.DBContext).DB

	pageNum, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	sortColumn := c.QueryParam("sort_column")

	// FIXME: опасно!
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY %s OFFSET $1 LIMIT $2", tableName, sortColumn)
	res, err := db.Query[T](DB, context.Background(), query, (pageNum-1)*pageSize, pageSize)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	query = fmt.Sprintf("SELECT count(*) FROM %s", tableName)

	var count int
	DB.QueryRow(context.Background(), query).Scan(&count)

	table := views.Table(listComp(res), pageNum, tableName, pageNum*pageSize >= count)
	return RenderPage(c, table)
}
