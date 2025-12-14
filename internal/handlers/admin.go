package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Timur1232/coursework_db/internal/db"
	"github.com/Timur1232/coursework_db/views"
	"github.com/a-h/templ"
	"github.com/jackc/pgx/v5"
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
	listComp, exist := TablesComponents[tableName]
	if !exist {
		return fmt.Errorf("table doesnt exist")
	}

	params := pgx.NamedArgs{
		"limit": pageSize,
	}

	query := fmt.Sprintf("SELECT * FROM %s LIMIT @limit", tableName)
	values, err := TableQueries[tableName](DB, context.Background(), query, params)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	query = fmt.Sprintf("SELECT count(*) FROM %s", tableName)
	var count int
	err = DB.QueryRow(context.Background(), query, params).Scan(&count)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fieldsComp := TablesFieldsComponents[tableName]

	isHxReq := c.Request().Header.Get("HX-Request") == "true"
	admin := views.AdminPanel(tableName, 1, pageSize, pageSize >= count, listComp(values), fieldsComp())
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

	query := fmt.Sprintf("SELECT * FROM %s", tableName)

	searchColumn := c.QueryParam("search")
	searchInput := c.FormValue("searchInput")
	if searchColumn != "" && searchInput != "" {
		query += fmt.Sprintf(" WHERE %s = @searchInput", searchColumn)
	}

	sortColumn := c.QueryParam("sortColumn")
	if sortColumn != "" {
		query += fmt.Sprintf(" ORDER BY %s", sortColumn)
	}

	query += " OFFSET @offset LIMIT @limit"
	params := pgx.NamedArgs{
		"searchInput": searchInput,
		"offset": (pageNum-1)*pageSize,
		"limit": pageSize,
	}

	values, err := TableQueries[tableName](DB, context.Background(), query, params)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	query = fmt.Sprintf("SELECT count(*) FROM %s", tableName)
	if searchColumn != "" && searchInput != "" {
		query += fmt.Sprintf(" WHERE %s = @searchInput", searchColumn)
	}

	var count int
	err = DB.QueryRow(context.Background(), query, params).Scan(&count)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	table := views.Table(listComp(values), pageNum, tableName, pageNum*pageSize >= count)
	return RenderPage(c, table)
}

// func AdminPanelEditRow(c echo.Context) error {
// 	DB := c.(*db.DBContext).DB
// 
// 	tableName := c.Param("tableName")
// 
// 	mainCol := c.QueryParam("col")
// 	mainVal := c.QueryParam("val")
// 	if mainCol == "" || mainVal == "" {
// 		return fmt.Errorf("no main field")
// 	}
// 
// 	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = $1", tableName, mainCol);
// 
// 	rows, err := TableQueries[tableName](DB, context.Background(), query, mainVal)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return err
// 	}
// 
// 	edit := Table(rows.([]any)[0])
// 
// 	return c.Redirect(http.StatusOK, fmt.Sprintf("/admin/%s", tableName))
// }

func AdminPanelDeleteRow(c echo.Context) error {
	DB := c.(*db.DBContext).DB

	tableName := c.Param("tableName")

	mainCol := c.QueryParam("col")
	mainVal := c.QueryParam("val")
	if mainCol == "" || mainVal == "" {
		return fmt.Errorf("no main field")
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE %s = $1", tableName, mainCol);

	_, err := TableQueries[tableName](DB, context.Background(), query, mainVal)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/admin/%s", tableName))
}
