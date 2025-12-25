package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func GetAllUsers(DB *pgx.Conn) ([]Users, error) {
	query := "SELECT * FROM users ORDER BY id_user"
	result, err := Query[Users](DB, context.Background(), query)
	if err != nil {
		return nil, err
	}
	return result.([]Users), nil
}

func AdminUpdateUserRole(DB *pgx.Conn, userID uint64, role Role) error {
	query := "UPDATE users SET role = $1 WHERE id_user = $2"
	_, err := DB.Exec(context.Background(), query, role, userID)
	return err
}

func GetUnlinkedApplications(DB *pgx.Conn) ([]ApplicationsForAdmission, error) {
	query := `
		SELECT * FROM applications_for_admission 
		WHERE id_user IS NULL OR id_user NOT IN (SELECT id_user FROM users)
		ORDER BY id_application
	`
	result, err := Query[ApplicationsForAdmission](DB, context.Background(), query)
	if err != nil {
		return nil, err
	}
	return result.([]ApplicationsForAdmission), nil
}

func GetUnlinkedRescuers(DB *pgx.Conn) ([]VgkRescuers, error) {
	query := `
		SELECT * FROM vgk_rescuers 
		WHERE id_user IS NULL OR id_user NOT IN (SELECT id_user FROM users)
		ORDER BY id_rescuer
	`
	result, err := Query[VgkRescuers](DB, context.Background(), query)
	if err != nil {
		return nil, err
	}
	return result.([]VgkRescuers), nil
}

func LinkApplicationToUser(DB *pgx.Conn, applicationID uint64, userID uint64) error {
	query := "UPDATE applications_for_admission SET id_user = $1 WHERE id_application = $2"
	_, err := DB.Exec(context.Background(), query, userID, applicationID)
	return err
}

func LinkRescuerToUser(DB *pgx.Conn, rescuerID uint64, userID uint64) error {
	query := "UPDATE vgk_rescuers SET id_user = $1 WHERE id_rescuer = $2"
	_, err := DB.Exec(context.Background(), query, userID, rescuerID)
	return err
}
