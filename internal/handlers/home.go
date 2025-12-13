package handlers

import (
	"context"
	"strconv"

	"github.com/Timur1232/coursework_db/internal/db"
	"github.com/Timur1232/coursework_db/views"
	"github.com/labstack/echo/v4"
)

func HomePage(c echo.Context) error {
	user := c.(*db.DBContext).User
	query := "SELECT * FROM accidents ORDER BY begin_date_time DESC LIMIT 10"
	accidents, err := db.Query[db.Accidents](c.(*db.DBContext).DB, context.Background(), query)
	if err != nil {
		return err
	}
	accidentsList := accidents.([]db.Accidents)

	home := views.HomePage(accidentsList)
	if c.Request().Header.Get("HX-Request") == "true" {
		return RenderPage(c, home)
	}
	page := views.Layout("Горноспасательная служба", home, user)
	return RenderPage(c, page)
}

func GetAccidents(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}
	pageSize := 10
	offset := (page - 1) * pageSize

	query := "SELECT * FROM accidents ORDER BY begin_date_time DESC LIMIT $1 OFFSET $2"
	accidents, err := db.Query[db.Accidents](c.(*db.DBContext).DB, context.Background(), query, pageSize, offset)
	if err != nil {
		return err
	}
	accidentsList := accidents.([]db.Accidents)

	hasNext := len(accidentsList) == pageSize

	return RenderPage(c, views.HomeAccidents(page, accidentsList, hasNext))
}

func GetObjects(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}
	pageSize := 10
	offset := (page - 1) * pageSize

	query := "SELECT * FROM objects ORDER BY name LIMIT $1 OFFSET $2"
	objects, err := db.Query[db.Objects](c.(*db.DBContext).DB, context.Background(), query, pageSize, offset)
	if err != nil {
		return err
	}
	objectsList := objects.([]db.Objects)

	hasNext := len(objectsList) == pageSize

	return RenderPage(c, views.HomeObjects(page, objectsList, hasNext))
}
