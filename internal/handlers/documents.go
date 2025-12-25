package handlers

import (
	"context"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Timur1232/coursework_db/internal/db"
	"github.com/Timur1232/coursework_db/views"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func ShowDocumentUploadForm(c echo.Context) error {
	user := c.(*db.DBContext).User

	if user.Role != db.Role_Candidate {
		return c.HTML(http.StatusOK, `<div class="message message-error">Недостаточно прав</div>`)
	}

	application, err := db.GetApplicationByUserID(c.(*db.DBContext).DB, user.IdUser)
	if err != nil {
		return c.HTML(http.StatusBadRequest, `<div class="message message-error">Заявление не найдено</div>`)
	}

	form := views.DocumentUploadForm(application.IdApplication, db.DocumentTypes)
	return RenderPage(c, form)
}

func CancelDocumentUpload(c echo.Context) error {
	return RenderPage(c, views.CancelUpload())
}

func UploadDocument(c echo.Context) error {
	user := c.(*db.DBContext).User

	if user.Role != db.Role_Candidate {
		return c.HTML(http.StatusOK, `<div class="message message-error">Недостаточно прав</div>`)
	}

	application, err := db.GetApplicationByUserID(c.(*db.DBContext).DB, user.IdUser)
	if err != nil {
		return c.HTML(http.StatusBadRequest, `<div class="message message-error">Заявление не найдено</div>`)
	}

	docType := c.FormValue("document_type")
	validUntilStr := c.FormValue("valid_until")

	if docType == "" || validUntilStr == "" {
		return c.HTML(http.StatusBadRequest, `<div class="message message-error">Заполните все поля</div>`)
	}

	validUntil, err := time.Parse("2006-01-02", validUntilStr)
	if err != nil {
		return c.HTML(http.StatusBadRequest, `<div class="message message-error">Некорректная дата</div>`)
	}

	file, err := c.FormFile("document_file")
	if err != nil {
		return c.HTML(http.StatusBadRequest, `<div class="message message-error">Выберите файл для загрузки</div>`)
	}

	if file.Size > 10*1024*1024 {
		return c.HTML(http.StatusBadRequest, `<div class="message message-error">Файл слишком большой (максимум 10MB)</div>`)
	}

	src, err := file.Open()
	if err != nil {
		return c.HTML(http.StatusOK, `<div class="message message-error">Ошибка при открытии файла</div>`)
	}
	defer src.Close()

	fileExt := filepath.Ext(file.Filename)
	fileName := uuid.New().String() + fileExt
	filePath := filepath.Join("static", "documents", fileName)

	if err := os.MkdirAll("static/documents", 0755); err != nil {
		return c.HTML(http.StatusOK, `<div class="message message-error">Ошибка при создании папки</div>`)
	}

	dst, err := os.Create(filePath)
	if err != nil {
		return c.HTML(http.StatusOK, `<div class="message message-error">Ошибка при создании файла</div>`)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {

		os.Remove(filePath)
		return c.HTML(http.StatusOK, `<div class="message message-error">Ошибка при сохранении файла</div>`)
	}

	documentURL := "/static/documents/" + fileName
	err = db.AddCandidateDocument(c.(*db.DBContext).DB, application.IdApplication, docType, documentURL, validUntil)
	if err != nil {

		os.Remove(filePath)
		return c.HTML(http.StatusOK, `<div class="message message-error">Ошибка при сохранении в базу данных</div>`)
	}

	return c.Redirect(http.StatusSeeOther, "/profile")
}

func DeleteDocument(c echo.Context) error {
	user := c.(*db.DBContext).User

	if user.Role != db.Role_Candidate {
		return c.HTML(http.StatusOK, `<div class="message message-error">Недостаточно прав</div>`)
	}

	applicationID, err := strconv.ParseUint(c.Param("application_id"), 10, 64)
	if err != nil {
		return c.HTML(http.StatusBadRequest, `<div class="message message-error">Некорректный ID заявления</div>`)
	}

	docType := c.Param("doc_type")

	application, err := db.GetApplicationByUserID(c.(*db.DBContext).DB, user.IdUser)
	if err != nil || application.IdApplication != applicationID {
		return c.HTML(http.StatusOK, `<div class="message message-error">Нет доступа к этому документу</div>`)
	}

	var documentURL string
	err = c.(*db.DBContext).DB.QueryRow(context.Background(),
		"SELECT document_url FROM candidates_documents WHERE id_application = $1 AND document_type = $2",
		applicationID, docType,
	).Scan(&documentURL)

	if err == nil && documentURL != "" {

		fileName := filepath.Base(documentURL)
		filePath := filepath.Join("static", "documents", fileName)

		os.Remove(filePath)
	}

	err = db.DeleteCandidateDocument(c.(*db.DBContext).DB, applicationID, docType)
	if err != nil {
		return c.HTML(http.StatusOK, `<div class="message message-error">Ошибка при удалении документа</div>`)
	}

	return c.Redirect(http.StatusSeeOther, "/profile")
}
