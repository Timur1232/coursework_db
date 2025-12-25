REVOKE CONNECT ON DATABASE coursework FROM PUBLIC;
REVOKE USAGE ON SCHEMA public FROM PUBLIC;

-- гость
GRANT CONNECT ON DATABASE coursework TO guest;
GRANT USAGE ON SCHEMA public TO guest;

GRANT SELECT ON TABLE objects TO guest;
GRANT SELECT ON TABLE accidents TO guest;
GRANT SELECT ON TABLE accident_types TO guest;

GRANT SELECT ON TABLE objects TO guest;
GRANT SELECT ON TABLE accidents TO guest;

-- кандидат
GRANT CONNECT ON DATABASE coursework TO candidate;
GRANT USAGE ON SCHEMA public TO candidate;
GRANT guest TO candidate;

GRANT SELECT, INSERT ON TABLE applications_for_admission TO candidate;
GRANT SELECT ON TABLE candidates_documents TO candidate;
GRANT SELECT ON TABLE candidates_medical_parameters TO candidate;
GRANT USAGE ON SEQUENCE applications_for_admission_id_application_seq TO candidate;

GRANT SELECT ON TABLE objects TO candidate;

-- член ВГК
GRANT CONNECT ON DATABASE coursework TO vgk_rescuer;
GRANT USAGE ON SCHEMA public TO vgk_rescuer;
GRANT guest TO vgk_rescuer;

GRANT SELECT ON TABLE vgk TO vgk_rescuer;
GRANT SELECT ON TABLE vgk_rescuers TO vgk_rescuer;
GRANT SELECT ON TABLE positions TO vgk_rescuer;
GRANT SELECT ON TABLE vgk_rescuers_documents TO vgk_rescuer;
GRANT SELECT ON TABLE vgk_locations TO vgk_rescuer;
GRANT SELECT ON TABLE vgk_shifts TO vgk_rescuer;
GRANT SELECT ON TABLE accidents_response_operations TO vgk_rescuer;
GRANT SELECT ON TABLE operations_participations TO vgk_rescuer;
GRANT SELECT ON TABLE operations_reports TO vgk_rescuer;
GRANT SELECT ON TABLE trainings TO vgk_rescuer;
GRANT SELECT ON TABLE trainings_participations TO vgk_rescuer;
GRANT SELECT ON TABLE certifications_passings TO vgk_rescuer;
GRANT SELECT ON TABLE vgk_rescuers_medical_parameters TO vgk_rescuer;
GRANT SELECT ON TABLE vgk_service_room TO vgk_rescuer;
GRANT SELECT ON TABLE equipment TO vgk_rescuer;
GRANT SELECT ON TABLE transport TO vgk_rescuer;
GRANT SELECT ON TABLE equipment_usage_history TO vgk_rescuer;
GRANT SELECT ON TABLE transport_usage_history TO vgk_rescuer;
GRANT SELECT ON TABLE equipment_service_history TO vgk_rescuer;
GRANT SELECT ON TABLE transport_service_history TO vgk_rescuer;
GRANT SELECT ON TABLE equipment_types TO vgk_rescuer;
GRANT SELECT ON TABLE accident_types TO vgk_rescuer;

-- оператор
GRANT CONNECT ON DATABASE coursework TO operator;
GRANT USAGE ON SCHEMA public TO operator;
GRANT vgk_rescuer TO operator;

GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO operator;
GRANT USAGE ON ALL SEQUENCES IN SCHEMA public TO operator;

GRANT EXECUTE ON FUNCTION transfer_application_to_rescuer(integer, integer) TO operator;

-- администратор
GRANT CONNECT ON DATABASE coursework TO admin;
GRANT USAGE ON SCHEMA public TO admin;
GRANT ALL PRIVILEGES ON SCHEMA public TO admin;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO admin;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO admin;
GRANT ALL PRIVILEGES ON ALL FUNCTIONS IN SCHEMA public TO admin;

-- приложение
GRANT CONNECT ON DATABASE coursework TO app;
GRANT USAGE ON SCHEMA public TO app;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO app;
GRANT USAGE ON ALL SEQUENCES IN SCHEMA public TO app;
GRANT EXECUTE ON ALL FUNCTIONS IN SCHEMA public TO app;
