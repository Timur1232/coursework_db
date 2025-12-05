SELECT document_type, id_application, document_url, valid_until FROM candidates_documents WHERE id_application = 1;

DELETE FROM applications_for_admission WHERE id_application = 1;

SELECT document_type, id_application, document_url, valid_until FROM candidates_documents WHERE id_application = 1;
