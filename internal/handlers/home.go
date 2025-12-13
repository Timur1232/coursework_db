package handlers

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func HomePage(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}

	pageSize := 10 // Количество элементов на странице
	offset := (page - 1) * pageSize

	// TODO:	
}
