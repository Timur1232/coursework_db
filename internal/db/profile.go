package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func GetApplicationByUserID(DB *pgx.Conn, userID uint64) (*ApplicationsForAdmission, error) {
	query := "SELECT * FROM applications_for_admission WHERE id_user = $1 LIMIT 1"
	result, err := Query[ApplicationsForAdmission](DB, context.Background(), query, userID)
	if err != nil {
		return nil, err
	}

	applications := result.([]ApplicationsForAdmission)
	if len(applications) == 0 {
		return nil, fmt.Errorf("application not found")
	}
	return &applications[0], nil
}

func GetCandidateDocuments(DB *pgx.Conn, applicationID uint64) ([]CandidatesDocuments, error) {
	query := "SELECT * FROM candidates_documents WHERE id_application = $1"
	result, err := Query[CandidatesDocuments](DB, context.Background(), query, applicationID)
	if err != nil {
		return nil, err
	}
	return result.([]CandidatesDocuments), nil
}

func GetCandidateMedicalParameters(DB *pgx.Conn, applicationID uint64) (*CandidatesMedicalParameters, error) {
	query := "SELECT * FROM candidates_medical_parameters WHERE id_application = $1 LIMIT 1"
	result, err := Query[CandidatesMedicalParameters](DB, context.Background(), query, applicationID)
	if err != nil {
		return nil, err
	}

	params := result.([]CandidatesMedicalParameters)
	if len(params) == 0 {
		return nil, fmt.Errorf("medical parameters not found")
	}
	return &params[0], nil
}

func GetRescuerByUserID(DB *pgx.Conn, userID uint64) (*VgkRescuers, error) {
	query := "SELECT * FROM vgk_rescuers WHERE id_user = $1 LIMIT 1"
	result, err := Query[VgkRescuers](DB, context.Background(), query, userID)
	if err != nil {
		return nil, err
	}

	rescuers := result.([]VgkRescuers)
	if len(rescuers) == 0 {
		return nil, fmt.Errorf("rescuer not found")
	}
	return &rescuers[0], nil
}

func GetRescuerDocuments(DB *pgx.Conn, rescuerID uint64) ([]VgkRescuersDocuments, error) {
	query := "SELECT * FROM vgk_rescuers_documents WHERE id_rescuer = $1"
	result, err := Query[VgkRescuersDocuments](DB, context.Background(), query, rescuerID)
	if err != nil {
		return nil, err
	}
	return result.([]VgkRescuersDocuments), nil
}

func GetRescuerMedicalParameters(DB *pgx.Conn, rescuerID uint64) ([]VgkRescuersMedicalParameters, error) {
	query := "SELECT * FROM vgk_rescuers_medical_parameters WHERE id_rescuer = $1 ORDER BY date DESC"
	result, err := Query[VgkRescuersMedicalParameters](DB, context.Background(), query, rescuerID)
	if err != nil {
		return nil, err
	}
	return result.([]VgkRescuersMedicalParameters), nil
}

func GetTeamMembers(DB *pgx.Conn, vgkID uint64) ([]VgkRescuers, error) {
	query := "SELECT * FROM vgk_rescuers WHERE id_vgk = $1 AND status != 'dismissed' ORDER BY position DESC NULLS LAST, first_name"
	result, err := Query[VgkRescuers](DB, context.Background(), query, vgkID)
	if err != nil {
		return nil, err
	}
	return result.([]VgkRescuers), nil
}

func GetVGKDetails(DB *pgx.Conn, vgkID uint64) (*Vgk, error) {
	query := "SELECT * FROM vgk WHERE id_vgk = $1 LIMIT 1"
	result, err := Query[Vgk](DB, context.Background(), query, vgkID)
	if err != nil {
		return nil, err
	}

	vgks := result.([]Vgk)
	if len(vgks) == 0 {
		return nil, fmt.Errorf("vgk not found")
	}
	return &vgks[0], nil
}

func DeleteApplication(DB *pgx.Conn, applicationID uint64) error {
	_, err := DB.Exec(context.Background(), "DELETE FROM applications_for_admission WHERE id_application = $1", applicationID)
	return err
}
