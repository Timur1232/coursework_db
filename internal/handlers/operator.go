package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Timur1232/coursework_db/internal/db"
	"github.com/Timur1232/coursework_db/views"
	"github.com/labstack/echo/v4"
)

func OperatorApplications(c echo.Context) error {
	user := c.(*db.DBContext).User

	applications, err := db.GetPendingApplications(c.(*db.DBContext).DB)
	if err != nil {
		fmt.Println(err)
		return err
	}

	applicationsList := views.OperatorApplicationsList(applications)
	if c.Request().Header.Get("HX-Request") == "true" {
		return RenderPage(c, applicationsList)
	}

	page := views.Layout("Заявления кандидатов", applicationsList, user)
	return RenderPage(c, page)
}

func OperatorApplicationDetail(c echo.Context) error {
	user := c.(*db.DBContext).User

	applicationID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/operator/applications")
	}

	application, documents, medicalParams, err := db.GetApplicationWithDetails(
		c.(*db.DBContext).DB, applicationID,
	)

	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/operator/applications")
	}

	detailView := views.OperatorApplicationDetail(application, documents, medicalParams)
	if c.Request().Header.Get("HX-Request") == "true" {
		return RenderPage(c, detailView)
	}

	page := views.Layout("Заявление кандидата", detailView, user)
	return RenderPage(c, page)
}

func ProcessApplication(c echo.Context) error {
	DB := c.(*db.DBContext).DB

	applicationID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.HTML(http.StatusOK, `<div class="message message-error">Некорректный ID заявления</div>`)
	}

	action := c.FormValue("action")

	switch action {
	case "approve":

		var candidateUserId sql.NullInt64
		DB.QueryRow(context.Background(), "SELECT id_user FROM applications_for_admission WHERE id_application = $1", applicationID).Scan(&candidateUserId)

		rows, err := DB.Query(context.Background(), "SELECT * FROM transfer_application_to_rescuer($1, $2)", applicationID, candidateUserId)
		if err != nil {
			fmt.Println(err)
			return c.HTML(http.StatusOK, `<div class="message message-error">Ошибка перевода</div>`)
		}
		var newRescuerId uint
		rows.Scan(&newRescuerId)
		rows.Close()

		if candidateUserId.Valid {
			_, err = DB.Exec(context.Background(), `UPDATE users SET role = 'rescuer' WHERE id_user = $1`, candidateUserId)
			if err != nil {
				fmt.Println(err)
				return c.HTML(http.StatusOK, `<div class="message message-error">Ошибка перевода</div>`)
			}
		}

	case "reject":

		err = db.UpdateApplicationStatus(c.(*db.DBContext).DB, applicationID, "отклонено")
		if err != nil {
			return c.HTML(http.StatusOK, `<div class="message message-error">Ошибка при обновлении статуса</div>`)
		}

	default:
		return c.HTML(http.StatusOK, `<div class="message message-error">Некорректное действие</div>`)
	}

	if c.Request().Header.Get("HX-Request") == "true" {
		return c.Redirect(http.StatusSeeOther, "/operator/applications")
	}

	return c.Redirect(http.StatusSeeOther, "/operator/applications")
}
