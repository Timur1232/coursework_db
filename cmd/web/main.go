package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Timur1232/coursework_db/internal/db"
	"github.com/Timur1232/coursework_db/internal/handlers"
	"github.com/Timur1232/coursework_db/views"

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

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*db.DBContext)

		cookie, err := c.Cookie("user_id")
		if err == nil && cookie != nil && cookie.Value != "" {
			userID, err := strconv.ParseUint(cookie.Value, 10, 64)
			if err == nil {
				user, err := db.GetUser(cc.DB, userID)
				if err == nil && user != nil {
					cc.User = user
				}
			}
		}

		return next(cc)
	}
}

func AuthRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*db.DBContext)
		if cc.User == nil {
			return c.Redirect(http.StatusSeeOther, "/login")
		}
		return next(c)
	}
}

func AuthRequiredAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*db.DBContext)
		if cc.User == nil {
			return c.Redirect(http.StatusSeeOther, "/login")
		} else if cc.User.Role != db.Role_Admin {
			return c.Redirect(http.StatusSeeOther, "/no_permission")
		}
		return next(c)
	}
}

func AuthRequiredOperator(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*db.DBContext)
		if cc.User == nil {
			return c.Redirect(http.StatusSeeOther, "/login")
		} else if cc.User.Role != db.Role_Operator {
			return c.Redirect(http.StatusSeeOther, "/no_permission")
		}
		return next(c)
	}
}

func AuthRequiredRescuer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*db.DBContext)
		if cc.User == nil {
			return c.Redirect(http.StatusSeeOther, "/login")
		} else if cc.User.Role != db.Role_Rescuer {
			return c.Redirect(http.StatusSeeOther, "/no_permission")
		}
		return next(c)
	}
}

func AuthRequiredCandidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*db.DBContext)
		if cc.User == nil {
			return c.Redirect(http.StatusSeeOther, "/login")
		} else if cc.User.Role != db.Role_Candidate {
			return c.Redirect(http.StatusSeeOther, "/no_permission")
		}
		return next(c)
	}
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
	e.Use(AuthMiddleware)
	e.Use(middleware.Logger())

	e.Static("/static", "static")

	adminGroup := e.Group("/admin", AuthRequiredAdmin)
	adminGroup.GET("", func(c echo.Context) error {
		return c.Redirect(http.StatusSeeOther, "/admin/vgk")
	})
	adminGroup.GET("/:tableName", handlers.AdminPanel)
	adminGroup.GET("/api/admin/:tableName/:page", handlers.AdminPanelPage)

	e.GET("/no_permission", func(c echo.Context) error {
		msg := views.NotAuthorizedNotification()
		if c.Request().Header.Get("HX-Request") == "true" {
			return handlers.RenderPage(c, msg)
		}
		page := views.Layout("Горноспасательная служба", msg, c.(*db.DBContext).User)
		return handlers.RenderPage(c, page)
	})

	e.GET("/", handlers.HomePage)
	e.GET("/login", handlers.Login)
	e.POST("/login", handlers.PostLogin)
	e.GET("/register", handlers.Register)
	e.POST("/register", handlers.PostRegister)
	e.POST("/logout", handlers.PostLogout)

	e.GET("/api/home/accidents", handlers.GetAccidents)
	e.GET("/api/home/objects", handlers.GetObjects)

	e.GET("/profile", handlers.Profile, AuthRequired)

	e.Logger.Fatal(e.Start(":42069"))

}
