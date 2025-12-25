package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

var DocumentTypes = []DocumentType{
	DocumentType_MedicalCertificate,
	DocumentType_EducationDiploma,
	DocumentType_TrainingCertificate,
	DocumentType_WorkBook,
	DocumentType_MilitaryId,
	DocumentType_PassportCopy,
	DocumentType_Photo,
}

func AddCandidateDocument(DB *pgx.Conn, applicationID uint64, docType string, documentURL string, validUntil time.Time) error {
	query := `
		INSERT INTO candidates_documents (document_type, id_application, document_url, valid_until)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (document_type, id_application)
		DO UPDATE SET document_url = $3, valid_until = $4
	`
	_, err := DB.Exec(context.Background(), query, docType, applicationID, documentURL, validUntil)
	return err
}

func DeleteCandidateDocument(DB *pgx.Conn, applicationID uint64, docType string) error {
	query := "DELETE FROM candidates_documents WHERE id_application = $1 AND document_type = $2"
	_, err := DB.Exec(context.Background(), query, applicationID, docType)
	return err
}
