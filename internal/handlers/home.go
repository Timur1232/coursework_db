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
	DB := c.(*db.DBContext).DB
	pageSize := 10

	query := "SELECT * FROM accidents ORDER BY begin_date_time DESC LIMIT $1"
	accidents, err := db.Query[db.Accidents](DB, context.Background(), query, pageSize)
	if err != nil {
		return err
	}
	accidentsList := accidents.([]db.Accidents)

	query = "SELECT COUNT(*) FROM accidents"
	var count int
	if err = DB.QueryRow(context.Background(), query).Scan(&count); err != nil {
		return err
	}

	hasNext := pageSize < count

	hxReq := c.Request().Header.Get("HX-Request") == "true"
	home := views.HomePage(accidentsList, hasNext, hxReq)
	if hxReq {
		return RenderPage(c, home)
	}
	page := views.Layout("Горноспасательная служба", home, user)
	return RenderPage(c, page)
}

func GetAccidents(c echo.Context) error {
	DB := c.(*db.DBContext).DB

	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}
	pageSize := 10
	offset := (page - 1) * pageSize

	query := "SELECT * FROM accidents ORDER BY begin_date_time DESC LIMIT $1 OFFSET $2"
	accidents, err := db.Query[db.Accidents](DB, context.Background(), query, pageSize, offset)
	if err != nil {
		return err
	}
	accidentsList := accidents.([]db.Accidents)

	query = "SELECT COUNT(*) FROM accidents"
	var count int
	if err = DB.QueryRow(context.Background(), query).Scan(&count); err != nil {
		return err
	}

	hasNext := page*pageSize < count

	if page == 1 {
		return RenderPage(c, views.HomeAccidents(page, accidentsList, hasNext))
	}

	return RenderPage(c, views.AccidentsRows(accidentsList, page, hasNext))
}

func GetObjects(c echo.Context) error {
	DB := c.(*db.DBContext).DB

	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}
	pageSize := 10
	offset := (page - 1) * pageSize

	query := "SELECT * FROM objects ORDER BY name LIMIT $1 OFFSET $2"
	objects, err := db.Query[db.Objects](DB, context.Background(), query, pageSize, offset)
	if err != nil {
		return err
	}
	objectsList := objects.([]db.Objects)

	query = "SELECT COUNT(*) FROM objects"
	var count int
	if err = DB.QueryRow(context.Background(), query).Scan(&count); err != nil {
		return err
	}

	hasNext := page*pageSize < count

	if page == 1 {
		return RenderPage(c, views.HomeObjects(page, objectsList, hasNext))
	}

	return RenderPage(c, views.ObjectsRows(objectsList, page, hasNext))
}
