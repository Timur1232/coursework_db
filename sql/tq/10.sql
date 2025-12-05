SELECT document_type, id_rescuer, document_url, valid_until FROM vgk_rescuers_documents WHERE id_rescuer = 1;

DELETE FROM vgk_rescuers WHERE id_rescuer = 1;

SELECT document_type, id_rescuer, document_url, valid_until FROM vgk_rescuers_documents WHERE id_rescuer = 1;
