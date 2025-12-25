package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Timur1232/coursework_db/internal/db"
	"github.com/Timur1232/coursework_db/views"
	"github.com/labstack/echo/v4"
)

func AdminUserManagement(c echo.Context) error {
	user := c.(*db.DBContext).User

	if user.Role != db.Role_Admin {
		return c.Redirect(http.StatusSeeOther, "/profile")
	}

	users, err := db.GetAllUsers(c.(*db.DBContext).DB)
	if err != nil {
		return err
	}

	managementView := views.AdminUserManagement(users)

	if c.Request().Header.Get("HX-Request") == "true" {
		return RenderPage(c, managementView)
	}

	page := views.Layout("Управление пользователями", managementView, user)
	return RenderPage(c, page)
}

func AdminAccountLinking(c echo.Context) error {
	user := c.(*db.DBContext).User

	if user.Role != db.Role_Admin {
		return c.Redirect(http.StatusSeeOther, "/profile")
	}

	unlinkedApplications, _ := db.GetUnlinkedApplications(c.(*db.DBContext).DB)

	unlinkedRescuers, _ := db.GetUnlinkedRescuers(c.(*db.DBContext).DB)

	allUsers, _ := db.GetAllUsers(c.(*db.DBContext).DB)

	linkingView := views.AdminAccountLinking(unlinkedApplications, unlinkedRescuers, allUsers)

	if c.Request().Header.Get("HX-Request") == "true" {
		return RenderPage(c, linkingView)
	}

	page := views.Layout("Привязка аккаунтов", linkingView, user)
	return RenderPage(c, page)
}

func UpdateUserRoleHandler(c echo.Context) error {
	user := c.(*db.DBContext).User

	if user.Role != db.Role_Admin {
		return c.String(http.StatusOK, "Недостаточно прав")
	}

	userID, err := strconv.ParseUint(c.FormValue("user_id"), 10, 64)
	if err != nil {
		return c.String(http.StatusOK, "Некорректный ID пользователя")
	}

	newRole := db.Role(c.FormValue("new_role"))

	validRoles := map[db.Role]bool{
		db.Role_Guest:     true,
		db.Role_Candidate: true,
		db.Role_Rescuer:   true,
		db.Role_Operator:  true,
		db.Role_Admin:     true,
	}

	if !validRoles[newRole] {
		fmt.Println(newRole)
		return c.String(http.StatusOK, "Некорректная роль")
	}

	err = db.ApplicationUpdateUserRole(c.(*db.DBContext).DB, userID, newRole)
	if err != nil {
		return c.String(http.StatusOK, "Ошибка при обновлении роли")
	}

	if c.Request().Header.Get("HX-Request") == "true" {

		updatedUser, err := db.GetUser(c.(*db.DBContext).DB, userID)
		if err != nil {
			return c.String(http.StatusOK, "Ошибка при получении данных пользователя")
		}

		return RenderPage(c, views.UserTableRow(updatedUser))
	}

	return AdminUserManagement(c)
}

func LinkApplicationHandler(c echo.Context) error {
	user := c.(*db.DBContext).User

	if user.Role != db.Role_Admin {
		return c.String(http.StatusOK, "Недостаточно прав")
	}

	applicationID, err := strconv.ParseUint(c.FormValue("application_id"), 10, 64)
	if err != nil {
		return c.String(http.StatusOK, "Некорректный ID заявления")
	}

	userID, err := strconv.ParseUint(c.FormValue("user_id"), 10, 64)
	if err != nil {
		return c.String(http.StatusOK, "Некорректный ID пользователя")
	}

	err = db.LinkApplicationToUser(c.(*db.DBContext).DB, applicationID, userID)
	if err != nil {
		return c.String(http.StatusOK, "Ошибка при привязке заявления")
	}

	if c.Request().Header.Get("HX-Request") == "true" {

		unlinkedApplications, _ := db.GetUnlinkedApplications(c.(*db.DBContext).DB)
		allUsers, _ := db.GetAllUsers(c.(*db.DBContext).DB)

		return RenderPage(c, views.UnlinkedApplicationsTable(unlinkedApplications, allUsers))
	}

	return AdminAccountLinking(c)
}

func LinkRescuerHandler(c echo.Context) error {
	user := c.(*db.DBContext).User

	if user.Role != db.Role_Admin {
		return c.String(http.StatusForbidden, "Недостаточно прав")
	}

	rescuerID, err := strconv.ParseUint(c.FormValue("rescuer_id"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Некорректный ID спасателя")
	}

	userID, err := strconv.ParseUint(c.FormValue("user_id"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Некорректный ID пользователя")
	}

	err = db.LinkRescuerToUser(c.(*db.DBContext).DB, rescuerID, userID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Ошибка при привязке спасателя")
	}

	if c.Request().Header.Get("HX-Request") == "true" {

		unlinkedRescuers, _ := db.GetUnlinkedRescuers(c.(*db.DBContext).DB)
		allUsers, _ := db.GetAllUsers(c.(*db.DBContext).DB)

		return RenderPage(c, views.UnlinkedRescuersTable(unlinkedRescuers, allUsers))
	}

	return AdminAccountLinking(c)
}
