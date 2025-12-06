package db

import (
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

type DBContext struct {
	echo.Context
	DB *pgx.Conn
}
