package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func GetAllObjects(DB *pgx.Conn) ([]Objects, error) {
	query := "SELECT * FROM objects ORDER BY name"
	result, err := Query[Objects](DB, context.Background(), query)
	if err != nil {
		return nil, err
	}
	return result.([]Objects), nil
}

func CreateApplication(DB *pgx.Conn, app *ApplicationsForAdmission) (uint64, error) {
	var id uint64
	query := `
		INSERT INTO applications_for_admission (
			id_object, passport_number, first_name, last_name, surname,
			issue_date, phone, email, status, birthday_date, home_address, id_user
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id_application
	`

	err := DB.QueryRow(context.Background(), query,
		app.IdObject, app.PassportNumber, app.FirstName, app.LastName, app.Surname,
		app.IssueDate, app.Phone, app.Email, app.Status, app.BirthdayDate, app.HomeAddress, app.IdUser.Int64,
	).Scan(&id)

	return id, err
}

func CreateCandidateDocument(DB *pgx.Conn, doc *CandidatesDocuments) error {
	query := `
		INSERT INTO candidates_documents (document_type, id_application, document_url, valid_until)
		VALUES ($1, $2, $3, $4)
	`
	_, err := DB.Exec(context.Background(), query,
		doc.DocumentType, doc.IdApplication, doc.DocumentUrl, doc.ValidUntil,
	)
	return err
}

func CreateCandidateMedicalParameters(DB *pgx.Conn, params *CandidatesMedicalParameters) error {
	query := `
		INSERT INTO candidates_medical_parameters 
		(id_application, date, expire_date, health_group, height, weight, note)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := DB.Exec(context.Background(), query,
		params.IdApplication, params.Date, params.ExpireDate,
		params.HealthGroup, params.Height, params.Weight, params.Note,
	)
	return err
}

func ApplicationUpdateUserRole(DB *pgx.Conn, userID uint64, role Role) error {
	query := "UPDATE users SET role = $1 WHERE id_user = $2"
	_, err := DB.Exec(context.Background(), query, role, userID)
	return err
}
