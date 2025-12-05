-- ========[enums]======== --

CREATE TYPE object_type_enum AS ENUM (
    'surface',
    'deep',
    'upland',
    'upland_deep',
    'underwater_mining'
);

CREATE TYPE danger_level_enum AS ENUM (
    'low',
    'medium',
    'high',
    'critical'
);

CREATE TYPE accident_status_enum AS ENUM (
    'reported',
    'in_progress',
    'contained',
    'resolved',
    'investigating'
);

CREATE TYPE document_type_enum AS ENUM (
    'medical_certificate',
    'education_diploma',
    'training_certificate',
    'work_book',
    'military_id',
    'passport_copy',
    'photo'
);

CREATE TYPE vgk_status_enum AS ENUM (
    'on_duty',
    'dismissed',
    'on_departure',
    'on_shift',
    'inactive'
);

CREATE TYPE vgk_location_status_enum AS ENUM (
    'operational',
    'malfunctioning'
);

CREATE TYPE operation_status_enum AS ENUM (
    'assessment',
    'evacuation',
    'aftermath_cleanup',
    'safety_ensuring',
    'completed',
    'failed'
);

CREATE TYPE equipment_status_enum AS ENUM (
    'operational',
    'in_use',
    'needs_repair_service',
    'under_repair',
    'written_off'
);

CREATE TYPE service_status_enum AS ENUM (
    'in_queue',
    'under_repair',
    'repaired',
    'irreparable'
);

-- ========[tables]======== --

CREATE TABLE IF NOT EXISTS equipment_types (
    type_name varchar(255) NOT NULL PRIMARY KEY,
    equipment_standards_url varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS objects (
    id_object integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    object_type object_type_enum NOT NULL,
    name varchar(255) NOT NULL,
    address varchar(255) NOT NULL,
    phone varchar(255) NOT NULL UNIQUE,
    email varchar(255) NOT NULL UNIQUE,
    director_full_name varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS accident_types (
    accident_name varchar(255) NOT NULL PRIMARY KEY,
    danger_level danger_level_enum NOT NULL
);

CREATE TABLE IF NOT EXISTS accidents (
    id_accident integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    id_object integer NOT NULL,
    accident_type varchar(255) NOT NULL,
    begin_date_time timestamp NOT NULL,
    status accident_status_enum NOT NULL,
    description text NOT NULL,
    first_estimate text NOT NULL,
    cause text NOT NULL
);

CREATE TABLE IF NOT EXISTS applications_for_admission (
    id_application integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    id_object integer NOT NULL,
    passport_number varchar(20) NOT NULL UNIQUE,
    first_name varchar(255) NOT NULL,
    last_name varchar(255) NOT NULL,
    surname varchar(255) DEFAULT NULL,
    issue_date date NOT NULL DEFAULT CURRENT_DATE,
    phone varchar(20) NOT NULL UNIQUE,
    email varchar(255) UNIQUE,
    status varchar(255) NOT NULL,
    birthday_date date NOT NULL,
    home_address varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS candidates_documents (
    document_type document_type_enum NOT NULL,
    id_application integer NOT NULL,
    document_url varchar(255) NOT NULL,
    valid_until date NOT NULL
);

CREATE TABLE IF NOT EXISTS candidates_medical_parameters (
    id_application integer PRIMARY KEY,
    date date NOT NULL,
    expire_date date NOT NULL,
    health_group integer NOT NULL,
    height decimal(5, 2) NOT NULL,
    weight decimal(5, 2) NOT NULL,
    note text NOT NULL
);

CREATE TABLE IF NOT EXISTS vgk (
    id_vgk integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    id_object integer NOT NULL,
    status vgk_status_enum NOT NULL DEFAULT 'inactive',
    formation_date date NOT NULL
);

CREATE TABLE IF NOT EXISTS positions (
    position_name varchar(255) NOT NULL PRIMARY KEY,
    salary decimal(10, 2) NOT NULL,
    min_experience_years integer DEFAULT 0,
    responsibilities text
);

CREATE TABLE IF NOT EXISTS vgk_rescuers (
    id_rescuer integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    id_vgk integer DEFAULT NULL,
    position varchar(255) DEFAULT NULL,
    first_name varchar(255) NOT NULL,
    second_name varchar(255) NOT NULL,
    surname varchar(255) DEFAULT NULL,
    status vgk_status_enum NOT NULL DEFAULT 'inactive',
    birth_date date NOT NULL,
    home_address varchar(255) NOT NULL,
    experience_years integer NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS vgk_rescuers_documents (
    document_type document_type_enum NOT NULL,
    id_rescuer integer NOT NULL,
    document_url varchar(255) NOT NULL,
    valid_until date NOT NULL
);

CREATE TABLE IF NOT EXISTS vgk_locations (
    id_vgk_location integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    id_responsible integer DEFAULT NULL,
    address varchar(255) NOT NULL,
    status vgk_location_status_enum NOT NULL
);

CREATE TABLE IF NOT EXISTS vgk_shifts (
    shift_start timestamp NOT NULL,
    id_vgk integer NOT NULL,
    id_vgk_location integer NOT NULL,
    shift_end timestamp NOT NULL
);

CREATE TABLE IF NOT EXISTS accidents_response_operations (
    id_operation integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    id_accident integer NOT NULL,
    start_date_time timestamp NOT NULL,
    end_date_time timestamp DEFAULT NULL,
    status operation_status_enum NOT NULL
);

CREATE TABLE IF NOT EXISTS operations_participations (
    id_vgk integer NOT NULL,
    id_operation integer NOT NULL,
    assigned_task varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS operations_reports (
    id_report integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    id_operation integer NOT NULL,
    report_date_time timestamp NOT NULL,
    description text NOT NULL
);

CREATE TABLE IF NOT EXISTS trainings (
    date date NOT NULL,
    id_object_location integer NOT NULL,
    id_instructor integer NOT NULL,
    topic varchar(255) NOT NULL,
    description text DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS trainings_participations (
    date date NOT NULL,
    id_object_location integer NOT NULL,
    id_rescuer integer NOT NULL,
    notes text DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS certifications_passings (
    date date NOT NULL,
    id_rescuer integer NOT NULL,
    result boolean NOT NULL,
    topic varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS vgk_rescuers_medical_parameters (
    date date NOT NULL,
    id_rescuer integer NOT NULL,
    expire_date date NOT NULL,
    health_group integer NOT NULL,
    height decimal NOT NULL,
    weight decimal NOT NULL,
    conclusion varchar(255) NOT NULL,
    note text DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS vgk_service_room (
    id_service_room integer GENERATED ALWAYS AS IDENTITY,
    id_responsible integer DEFAULT NULL,
    purpose varchar(255) NOT NULL,
    address varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS equipment (
    inventory_number integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    id_vgk_location integer DEFAULT NULL,
    equipment_type varchar(255) NOT NULL,
    name varchar(255) NOT NULL,
    status equipment_status_enum NOT NULL,
    last_check_date date NOT NULL
);

CREATE TABLE IF NOT EXISTS transport (
    transport_number integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    id_vgk_location integer DEFAULT NULL,
    model varchar(255) NOT NULL,
    type varchar(255) NOT NULL,
    status equipment_status_enum NOT NULL,
    manufacture_date date NOT NULL,
    mileage decimal NOT NULL,
    last_check_date date NOT NULL
);

CREATE TABLE IF NOT EXISTS equipment_usage_history (
    inventory_number integer NOT NULL,
    id_rescuer integer NOT NULL,
    issue_date date NOT NULL,
    return_date date NOT NULL,
    purpose varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS transport_usage_history (
    transport_number integer NOT NULL,
    id_rescuer integer NOT NULL,
    departure_date date NOT NULL,
    return_date date NOT NULL,
    purpose varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS equipment_service_history (
    inventory_number integer NOT NULL,
    id_service_room integer NOT NULL,
    reason varchar(255) NOT NULL,
    serve_date date NOT NULL,
    status service_status_enum NOT NULL
);

CREATE TABLE IF NOT EXISTS transport_service_history (
    transport_number integer NOT NULL,
    id_service_room integer NOT NULL,
    reason varchar(255) NOT NULL,
    serve_date date NOT NULL,
    status service_status_enum NOT NULL
);
