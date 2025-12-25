package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func GetPendingApplications(DB *pgx.Conn) ([]ApplicationsForAdmission, error) {
	query := `
		SELECT a.*
		FROM applications_for_admission a
		WHERE a.status = 'рассмотрение'
		ORDER BY a.issue_date DESC
	`

	result, err := Query[ApplicationsForAdmission](DB, context.Background(), query)
	if err != nil {
		return nil, err
	}
	return result.([]ApplicationsForAdmission), nil
}

func GetApplicationWithDetails(DB *pgx.Conn, applicationID uint64) (*ApplicationsForAdmission, []CandidatesDocuments, *CandidatesMedicalParameters, error) {
	query := `
		SELECT a.*
		FROM applications_for_admission a
		WHERE a.id_application = $1
	`

	result, err := Query[ApplicationsForAdmission](DB, context.Background(), query, applicationID)
	if err != nil {
		return nil, nil, nil, err
	}

	applications := result.([]ApplicationsForAdmission)
	if len(applications) == 0 {
		return nil, nil, nil, fmt.Errorf("application not found")
	}
	application := &applications[0]

	documents, _ := GetCandidateDocuments(DB, applicationID)

	medicalParams, _ := GetCandidateMedicalParameters(DB, applicationID)

	return application, documents, medicalParams, nil
}

func UpdateApplicationStatus(DB *pgx.Conn, applicationID uint64, status string) error {
	query := "UPDATE applications_for_admission SET status = $1 WHERE id_application = $2"
	_, err := DB.Exec(context.Background(), query, status, applicationID)
	return err
}

func TransferApplicationToRescuer(DB *pgx.Conn, applicationID uint64) error {

	application, documents, medicalParams, err := GetApplicationWithDetails(DB, applicationID)
	if err != nil {
		return err
	}

	var vgkID uint64
	err = DB.QueryRow(context.Background(),
		"SELECT id_vgk FROM vgk WHERE id_object = $1 LIMIT 1",
		application.IdObject,
	).Scan(&vgkID)

	if err != nil {
		err = DB.QueryRow(context.Background(),
			"INSERT INTO vgk (id_object, status, formation_date) VALUES ($1, $2, CURRENT_DATE) RETURNING id_vgk",
			application.IdObject, VgkStatus_Inactive,
		).Scan(&vgkID)
		if err != nil {
			return err
		}
	}

	var rescuerID uint64
	err = DB.QueryRow(context.Background(), `
		INSERT INTO vgk_rescuers (
			id_vgk, position, first_name, second_name, surname, 
			status, birth_date, home_address, experience_years, id_user
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id_rescuer
	`,
		vgkID,
		nil,
		application.FirstName,
		application.LastName,
		application.Surname,
		VgkStatus_Inactive,
		application.BirthdayDate,
		application.HomeAddress,
		0,
		application.IdUser.Int64,
	).Scan(&rescuerID)

	if err != nil {
		return err
	}

	if medicalParams != nil {
		_, err = DB.Exec(context.Background(), `
			INSERT INTO vgk_rescuers_medical_parameters (
				date, id_rescuer, expire_date, health_group, height, weight, conclusion, note
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		`,
			medicalParams.Date,
			rescuerID,
			medicalParams.ExpireDate,
			medicalParams.HealthGroup,
			medicalParams.Height,
			medicalParams.Weight,
			"Принят на основании заявления",
			medicalParams.Note,
		)
		if err != nil {
			return err
		}
	}

	for _, doc := range documents {
		_, err = DB.Exec(context.Background(), `
			INSERT INTO vgk_rescuers_documents (
				document_type, id_rescuer, document_url, valid_until
			) VALUES ($1, $2, $3, $4)
		`,
			doc.DocumentType,
			rescuerID,
			doc.DocumentUrl,
			doc.ValidUntil,
		)
		if err != nil {
			return err
		}
	}

	_, err = DB.Exec(context.Background(),
		"UPDATE users SET role = $1 WHERE id_user = $2",
		Role_Rescuer, application.IdUser.Int64,
	)

	return err
}
