package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Timur1232/coursework_db/internal/db"
	"github.com/Timur1232/coursework_db/views"
	"github.com/labstack/echo/v4"
)

func ShowApplicationForm(c echo.Context) error {
	user := c.(*db.DBContext).User

	if user.Role != db.Role_Guest {
		return c.Redirect(http.StatusSeeOther, "/profile")
	}

	objects, err := db.GetAllObjects(c.(*db.DBContext).DB)
	if err != nil {
		return err
	}

	appForm := views.ApplicationForm(objects)
	if c.Request().Header.Get("HX-Request") == "true" {
		return RenderPage(c, appForm)
	}
	page := views.Layout("Подача заявления", appForm, user)
	return RenderPage(c, page)
}

func SubmitApplication(c echo.Context) error {
	user := c.(*db.DBContext).User

	if user.Role != db.Role_Guest {
		return c.HTML(http.StatusBadRequest, `<div class="message message-error">Вы уже имеете статус выше гостя</div>`)
	}

	if err := c.Request().ParseForm(); err != nil {
		return c.HTML(http.StatusBadRequest, `<div class="message message-error">Ошибка обработки формы</div>`)
	}

	lastName := c.FormValue("last_name")
	firstName := c.FormValue("first_name")
	surname := c.FormValue("surname")
	passportNumber := c.FormValue("passport_number")
	phone := c.FormValue("phone")
	email := c.FormValue("email")
	homeAddress := c.FormValue("home_address")

	var count int
	err := c.(*db.DBContext).DB.QueryRow(context.Background(),
		"SELECT COUNT(*) FROM applications_for_admission WHERE passport_number = $1", passportNumber,
	).Scan(&count)
	if err != nil || count > 0 {
		return c.HTML(http.StatusOK, `<div class="message message-error">Заявление с таким номером паспорта уже существует</div>`)
	}

	err = c.(*db.DBContext).DB.QueryRow(context.Background(),
		"SELECT COUNT(*) FROM applications_for_admission WHERE phone = $1", phone,
	).Scan(&count)
	if err != nil || count > 0 {
		return c.HTML(http.StatusOK, `<div class="message message-error">Заявление с таким телефоном уже существует</div>`)
	}

	err = c.(*db.DBContext).DB.QueryRow(context.Background(),
		"SELECT COUNT(*) FROM applications_for_admission WHERE email = $1", email,
	).Scan(&count)
	if err != nil || count > 0 {
		return c.HTML(http.StatusOK, `<div class="message message-error">Заявление с таким email уже существует</div>`)
	}

	birthdayDate, _ := time.Parse("2006-01-02", c.FormValue("birthday_date"))
	issueDate := time.Now()

	idObject, _ := strconv.ParseUint(c.FormValue("id_object"), 10, 64)

	application := &db.ApplicationsForAdmission{
		IdObject:       idObject,
		PassportNumber: passportNumber,
		FirstName:      firstName,
		LastName:       lastName,
		Surname:        sql.NullString{String: surname, Valid: surname != ""},
		IssueDate:      issueDate,
		Phone:          phone,
		Email:          email,
		Status:         "рассмотрение",
		BirthdayDate:   birthdayDate,
		HomeAddress:    homeAddress,
		IdUser:         sql.NullInt64{Int64: int64(user.IdUser), Valid: true},
	}

	appID, err := db.CreateApplication(c.(*db.DBContext).DB, application)
	if err != nil {
		fmt.Println(err)
		return c.HTML(http.StatusOK, `<div class="message message-error">Ошибка при создании заявления</div>`)
	}

	medicalDate, _ := time.Parse("2006-01-02", c.FormValue("medical_date"))
	expireDate, _ := time.Parse("2006-01-02", c.FormValue("expire_date"))
	healthGroup := c.FormValue("health_group")
	height, _ := strconv.ParseFloat(c.FormValue("height"), 32)
	weight, _ := strconv.ParseFloat(c.FormValue("weight"), 32)
	medicalNote := c.FormValue("medical_note")

	medicalParams := &db.CandidatesMedicalParameters{
		IdApplication: appID,
		Date:          medicalDate,
		ExpireDate:    expireDate,
		HealthGroup:   healthGroup,
		Height:        float32(height),
		Weight:        float32(weight),
		Note:          medicalNote,
	}

	if err := db.CreateCandidateMedicalParameters(c.(*db.DBContext).DB, medicalParams); err != nil {
		fmt.Println(err)
		return c.HTML(http.StatusOK, `<div class="message message-error">Ошибка при сохранении медицинских параметров</div>`)
	}

	if err := db.ApplicationUpdateUserRole(c.(*db.DBContext).DB, user.IdUser, db.Role_Candidate); err != nil {
		fmt.Println(err)
		return c.HTML(http.StatusOK, `<div class="message message-error">Ошибка при обновлении роли</div>`)
	}

	user.Role = db.Role_Candidate

	if c.Request().Header.Get("HX-Request") == "true" {
		c.Response().Header().Set("HX-Redirect", "/profile")
		return c.NoContent(http.StatusOK)
	}

	successPage := views.ApplicationSuccess()
	page := views.Layout("Заявление подано", successPage, user)
	return RenderPage(c, page)
}
