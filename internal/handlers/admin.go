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

func AdminPanel(c echo.Context) error {
	DB := c.(*db.DBContext).DB

	tableName := c.Param("tableName")
	sortComp, exist := TablesSortComponents[tableName]
	if !exist {
		return fmt.Errorf("table %s doesnt exist", tableName)
	}
	listComp, exist := TablesComponents[tableName]
	if !exist {
		return fmt.Errorf("table doesnt exist")
	}

	query := fmt.Sprintf("SELECT * FROM %s LIMIT $1", tableName)
	values, err := TableQueries[tableName](DB, context.Background(), query, pageSize)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	query = fmt.Sprintf("SELECT count(*) FROM %s", tableName)
	var count int
	err = DB.QueryRow(context.Background(), query).Scan(&count)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	isHxReq := c.Request().Header.Get("HX-Request") == "true"
	admin := views.AdminPanel(tableName, 1, pageSize, pageSize >= count, listComp(values), sortComp())
	if isHxReq {
		return RenderPage(c, admin)
	}

	page := views.Layout("Горноспасательная служба", admin, &db.Users{IdUser: 69, Login: "Timur Baimuradov", Password: "a$$word", Role: db.Role_Admin})
	return RenderPage(c, page)
}

func AdminPanelPage(c echo.Context) error {
	DB := c.(*db.DBContext).DB

	pageNum, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	tableName := c.Param("tableName")

	listComp, exist := TablesComponents[tableName]
	if !exist {
		return fmt.Errorf("table %s doesnt exist", tableName)
	}

	sortColumn := c.QueryParam("sortColumn")
	fmt.Println("sortColumn =", sortColumn)
	var query string
	if sortColumn == "" {
		query = fmt.Sprintf("SELECT * FROM %s OFFSET $1 LIMIT $2", tableName)
	} else {
		query = fmt.Sprintf("SELECT * FROM %s ORDER BY %s OFFSET $1 LIMIT $2", tableName, sortColumn)
	}

	values, err := TableQueries[tableName](DB, context.Background(), query, (pageNum-1)*pageSize, pageSize)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	query = fmt.Sprintf("SELECT count(*) FROM %s", tableName)
	var count int
	err = DB.QueryRow(context.Background(), query).Scan(&count)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	table := views.Table(listComp(values), pageNum, tableName, pageNum*pageSize >= count)
	return RenderPage(c, table)
}
