CREATE TABLE equipment_types (
    type_name varchar(255) NOT NULL PRIMARY KEY,
    equipment_standards_url varchar(255) NOT NULL
);

CREATE TYPE object_type_enum AS ENUM (
    'surface',
    'deep',
    'upland',
    'upland_deep',
    'underwater_mining'
);

CREATE TABLE objects (
    id_object integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    object_type object_type_enum NOT NULL,
    name varchar(255) NOT NULL,
    address varchar(255) NOT NULL,
    phone varchar(255) NOT NULL UNIQUE,
    email varchar(255) NOT NULL UNIQUE,
    director_full_name varchar(255) NOT NULL,

    CONSTRAINT phone_format_check CHECK (
        phone ~ '^\+7\(\d{3}\)\d{3}-\d{2}-\d{2}$'
    ),

    CONSTRAINT email_format_check CHECK (
        email ~ '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'
    )
);

CREATE TYPE danger_level_enum AS ENUM (
    'low',
    'medium',
    'high',
    'critical'
);

CREATE TABLE accident_types (
    accident_name varchar(255) NOT NULL PRIMARY KEY,
    danger_level danger_level_enum NOT NULL
);

CREATE TYPE accident_status_enum AS ENUM (
    'reported',
    'in_progress',
    'contained',
    'resolved',
    'investigating'
);

CREATE TABLE accidents (
    id_accident integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    id_object integer NOT NULL,
    accident_type varchar(255) NOT NULL,
    begin_date_time timestamp NOT NULL,
    status accident_status_enum NOT NULL,
    description text NOT NULL,
    first_estimate text NOT NULL,
    cause text NOT NULL,

    CONSTRAINT accidents_id_object_objects_id_object_foreign FOREIGN KEY (id_object)
        REFERENCES objects (id_object)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,

    CONSTRAINT accidents_id_accident_type_foreign FOREIGN KEY (accident_type)
        REFERENCES accident_types (accident_name)
        ON UPDATE CASCADE
        ON DELETE RESTRICT
);

CREATE TABLE applications_for_admission (
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

    CONSTRAINT applications_for_admission_id_object_objects_id_object_foreign FOREIGN KEY (id_object) REFERENCES objects (id_object)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,

    CONSTRAINT phone_format_check CHECK (
        phone ~ '^\+7\(\d{3}\)\d{3}-\d{2}-\d{2}$'
    ),

    CONSTRAINT email_format_check CHECK (
        email IS NULL OR email ~ '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$'
    ),

    CONSTRAINT age_check CHECK (
        birthday_date <= CURRENT_DATE - INTERVAL '18 years'
    )
);

CREATE TYPE document_type_enum AS ENUM (
    'medical_certificate',
    'education_diploma',
    'training_certificate',
    'work_experience',
    'military_id',
    'passport_copy',
    'photo'
);

CREATE TABLE candidates_documents (
    document_type document_type_enum NOT NULL,
    id_application integer NOT NULL,
    document_url varchar(255) NOT NULL,
    valid_until date NOT NULL,

    PRIMARY KEY (document_type, id_application),

    CONSTRAINT valid_until_check CHECK (
        valid_until >= CURRENT_DATE
    ),

    CONSTRAINT candidates_documents_id_application_applications_for_admission_id_application_foreign FOREIGN KEY (id_application) REFERENCES applications_for_admission (id_application)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE candidates_medical_parameters (
    id_application integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    date date NOT NULL,
    expire_date date NOT NULL,
    health_group integer NOT NULL,
    height decimal(5, 2) NOT NULL,
    weight decimal(5, 2) NOT NULL,
    note text NOT NULL,

    CONSTRAINT date_not_future_check CHECK (date <= CURRENT_DATE),
    CONSTRAINT expire_after_date_check CHECK (expire_date > date),
    CONSTRAINT expire_date_check CHECK (expire_date >= CURRENT_DATE),

    CONSTRAINT height_check CHECK (height BETWEEN 0.00 AND 350.00),
    CONSTRAINT weight_check CHECK (weight BETWEEN 0.00 AND 250.00),

    CONSTRAINT candidates_medical_parameters_id_application_applications_for_admission_id_application_foreign FOREIGN KEY (id_application) REFERENCES applications_for_admission (id_application)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TYPE vgk_status_enum AS ENUM (
    'disbanded',
    'on_departure',
    'on_shift',
    'temporarily_inactive'
);

CREATE TABLE vgk (
    id_vgk integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    id_object integer NOT NULL,
    status vgk_status_enum NOT NULL,
    formation_date date NOT NULL,

    CONSTRAINT formation_date_check CHECK (formation_date <= current_date),

    CONSTRAINT vgk_id_object_objects_id_object_foreign FOREIGN KEY (id_object) REFERENCES objects (id_object)
        ON UPDATE CASCADE
        ON DELETE RESTRICT
);

CREATE TABLE positions (
    position_name varchar(255) NOT NULL PRIMARY KEY,
    salary decimal(10, 2) NOT NULL,
    min_experience_years integer DEFAULT 0,
    responsibilities text,

    CONSTRAINT salary_check CHECK (
        salary BETWEEN 16242.00 AND 1000000.00 -- мрот
    ),

    CONSTRAINT min_experience_check CHECK (
        min_experience_years BETWEEN 0 AND 50
    )
);

CREATE TYPE rescuer_status_enum AS ENUM (
    'on_duty',
    'on_departure',
    'dismissed',
    'on_vacation'
);

CREATE TABLE vgk_rescuers (
    id_rescuer integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    id_vgk integer DEFAULT NULL,
    position varchar(255) DEFAULT NULL,
    first_name varchar(255) NOT NULL,
    second_name varchar(255) NOT NULL,
    surname varchar(255) DEFAULT NULL,
    status rescuer_status_enum NOT NULL,
    birth_date date NOT NULL,
    home_address varchar(255) NOT NULL,
    experience_years integer NOT NULL DEFAULT 0,

    CONSTRAINT birth_date_check CHECK (
        birth_date <= current_date - interval '18 years'
    ),

    CONSTRAINT experience_check CHECK (
        experience_years BETWEEN 0 AND 50
    ),

    CONSTRAINT commander_must_be_in_vgk CHECK (
        NOT (position = 'commander' AND id_vgk IS NULL)
    ),

    CONSTRAINT dismissed_not_in_vgk CHECK (
        NOT (status = 'dismissed' AND id_vgk IS NOT NULL)
    ),

    CONSTRAINT vgk_rescuers_id_vgk_vgk_id_vgk_foreign FOREIGN KEY (id_vgk) REFERENCES vgk (id_vgk)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,

    CONSTRAINT position_foreign FOREIGN KEY (position) REFERENCES positions (position_name)
        ON UPDATE CASCADE
        ON DELETE RESTRICT
);

CREATE TABLE vgk_rescuers_documents (
    document_type document_type_enum NOT NULL,
    id_rescuer integer NOT NULL,
    document_url varchar(255) NOT NULL,
    valid_until date NOT NULL,

    PRIMARY KEY (document_type, id_rescuer),

    CONSTRAINT valid_until_check CHECK (
        valid_until >= current_date
    ),

    CONSTRAINT vgk_rescuers_documents_id_rescuer_vgk_rescuers_id_rescuer_foreign FOREIGN KEY (id_rescuer) REFERENCES vgk_rescuers (id_rescuer)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE vgk_shifts (
    shift_start timestamp NOT NULL,
    id_vgk integer NOT NULL,
    id_vgk_location integer NOT NULL,
    shift_end timestamp NOT NULL,

    PRIMARY KEY (
        shift_start, id_vgk, id_vgk_location
    ),

    CONSTRAINT shift_end_after_start_check CHECK (
        shift_end > shift_start
    ),

    CONSTRAINT shift_start_not_future_check CHECK (
        shift_start <= current_timestamp
    ),

    CONSTRAINT vgk_shifts_id_vgk_vgk_id_vgk_foreign FOREIGN KEY (id_vgk) REFERENCES vgk (id_vgk)
        ON UPDATE CASCADE
        ON DELETE CASCADE,

    CONSTRAINT vgk_shifts_id_vgk_location_vgk_locations_id_vgk_location_foreign FOREIGN KEY (id_vgk_location) REFERENCES vgk_locations (id_vgk_location)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TYPE operation_status_enum AS ENUM (
    'assessment',
    'evacuation',
    'aftermath_cleanup',
    'safety_ensuring',
    'completed',
    'failed'
);

CREATE TABLE accidents_response_operations (
    id_operation integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    id_accident integer NOT NULL,
    id_leader integer NOT NULL,
    start_date_time timestamp NOT NULL,
    end_date_time timestamp DEFAULT NULL,
    status operation_status_enum NOT NULL,

    CONSTRAINT start_date_time_check CHECK (
        start_date_time <= current_timestamp
    ),

    CONSTRAINT end_date_time_check CHECK (
        end_date_time IS NULL OR end_date_time > start_date_time
    ),

    CONSTRAINT completed_has_end_date CHECK (
        NOT (status IN ('completed', 'failed') AND end_date_time IS NULL)
    ),

    CONSTRAINT accidents_response_operations_id_leader_vgk_rescuers_id_rescuer_foreign FOREIGN KEY (id_leader) REFERENCES vgk_rescuers (id_rescuer)
        ON UPDATE CASCADE
        ON DELETE CASCADE,

    CONSTRAINT accidents_response_operations_id_accident_accidents_id_accident_foreign FOREIGN KEY (id_accident) REFERENCES accidents (id_accident)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE operations_participations (
    id_vgk integer NOT NULL,
    id_operation integer NOT NULL,
    assigned_task varchar(255) NOT NULL,

    PRIMARY KEY (id_vgk, id_operation),

    CONSTRAINT assigned_task_not_empty_check CHECK (
        length(trim(assigned_task)) > 0
    ),

    CONSTRAINT operations_participations_id_vgk_vgk_id_vgk_foreign FOREIGN KEY (id_vgk) REFERENCES vgk (id_vgk)
        ON UPDATE CASCADE
        ON DELETE CASCADE,

    CONSTRAINT operations_participations_id_operation_accidents_response_operations_id_operation_foreign FOREIGN KEY (id_operation) REFERENCES accidents_response_operations (id_operation)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE operations_reports (
    id_report integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    id_operation integer NOT NULL,
    report_date_time timestamp NOT NULL,
    description text NOT NULL,

    CONSTRAINT report_date_time_check CHECK (
        report_date_time <= current_timestamp
    ),

    CONSTRAINT description_not_empty_check CHECK (
        length(trim(description)) > 0
    ),

    CONSTRAINT operations_reports_id_participation_accidents_response_operations_id_operation_foreign FOREIGN KEY (id_operation) REFERENCES accidents_response_operations (id_operation)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE trainings (
    date date NOT NULL,
    id_object_location integer NOT NULL,
    id_instructor integer NOT NULL,
    topic varchar(255) NOT NULL,
    description text DEFAULT NULL,

    PRIMARY KEY (date, id_object_location),

    CONSTRAINT date_check CHECK (
        date <= current_date
    ),

    CONSTRAINT trainings_id_object_location_objects_id_object_foreign FOREIGN KEY (id_object_location) REFERENCES objects (id_object)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,

    CONSTRAINT trainings_id_instructor_vgk_rescuers_id_rescuer_foreign FOREIGN KEY (id_instructor) REFERENCES vgk_rescuers (id_rescuer)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE trainings_participations (
    date date NOT NULL,
    id_rescuer integer NOT NULL,
    notes text DEFAULT NULL,

    PRIMARY KEY (date, id_rescuer),

    CONSTRAINT date_check CHECK (
        date <= current_date
    ),

    CONSTRAINT trainings_participations_id_rescuer_vgk_rescuers_id_rescuer_foreign FOREIGN KEY (id_rescuer) REFERENCES vgk_rescuers (id_rescuer)
        ON UPDATE CASCADE
        ON DELETE CASCADE,

    CONSTRAINT trainings_participations_date_trainings_date_foreign FOREIGN KEY (date) REFERENCES trainings (date)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE certifications_passings (
    date date NOT NULL,
    id_rescuer integer NOT NULL,
    result boolean NOT NULL,
    topic varchar(255) NOT NULL,

    PRIMARY KEY (date, id_rescuer),

    CONSTRAINT date_check CHECK (
        date <= current_date
    ),

    CONSTRAINT certifications_passings_id_rescuer_vgk_rescuers_id_rescuer_foreign FOREIGN KEY (id_rescuer) REFERENCES vgk_rescuers (id_rescuer)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE vgk_rescuers_medical_parameters (
    date date NOT NULL,
    id_rescuer integer NOT NULL,
    expire_date date NOT NULL,
    health_group integer NOT NULL,
    height decimal NOT NULL,
    weight decimal NOT NULL,
    conclusion varchar(255) NOT NULL,
    note text DEFAULT NULL,

    PRIMARY KEY (date, id_rescuer),

    CONSTRAINT date_check CHECK (date <= current_date),
    CONSTRAINT expire_date_check CHECK (
        expire_date > date AND expire_date >= current_date
    ),

    CONSTRAINT health_group_check CHECK (
        health_group BETWEEN 1 AND 5
    ),

    CONSTRAINT height_check CHECK (height BETWEEN 0.00 AND 350.00),
    CONSTRAINT weight_check CHECK (weight BETWEEN 0.00 AND 250.00),

    CONSTRAINT conclusion_not_empty_check CHECK (
        length(trim(conclusion)) > 0
    ),

    CONSTRAINT vgk_rescuers_medical_parameters_id_rescuer_vgk_rescuers_id_rescuer_foreign FOREIGN KEY (id_rescuer) REFERENCES vgk_rescuers (id_rescuer)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE vgk_service_room (
    id_service_room integer NOT NULL,
    id_responsible integer DEFAULT NULL,
    purpose varchar(255) NOT NULL,
    address varchar(255) NOT NULL,

    PRIMARY KEY (id_service_room),

    CONSTRAINT purpose_not_empty_check CHECK (
        length(trim(purpose)) > 0
    ),

    CONSTRAINT address_not_empty_check CHECK (
        length(trim(address)) > 0
    ),

    CONSTRAINT vgk_service_room_id_responsible_vgk_rescuers_id_rescuer_foreign FOREIGN KEY (id_responsible) REFERENCES vgk_rescuers (id_rescuer)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TYPE vgk_location_status_enum AS ENUM (
    'operational',
    'malfunctioning'
);

CREATE TABLE vgk_locations (
    id_vgk_location integer NOT NULL,
    id_responsible integer DEFAULT NULL,
    address varchar(255) NOT NULL,
    status vgk_location_status_enum NOT NULL,

    PRIMARY KEY (id_vgk_location),

    CONSTRAINT vgk_locations_id_responsible_vgk_rescuers_id_rescuer_foreign FOREIGN KEY (id_responsible) REFERENCES vgk_rescuers (id_rescuer)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TYPE equipment_status_enum AS ENUM (
    'operational',
    'in_use',
    'under_repair',
    'written_off'
);

CREATE TABLE equipment (
    inventory_number integer GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    id_vgk_location integer DEFAULT NULL,
    equipment_type varchar(255) NOT NULL,
    name varchar(255) NOT NULL,
    status equipment_status_enum NOT NULL,
    last_check_date date NOT NULL,

    CONSTRAINT last_check_date_check CHECK (
        last_check_date <= current_date
    ),

    CONSTRAINT equipment_id_vgk_location_vgk_locations_id_vgk_location_foreign FOREIGN KEY (id_vgk_location) REFERENCES vgk_locations (id_vgk_location)
        ON UPDATE CASCADE
        ON DELETE CASCADE,

    CONSTRAINT equipment_equipment_type_equipment_types_type_name_foreign FOREIGN KEY (equipment_type) REFERENCES equipment_types (type_name)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TYPE transport_status_enum AS ENUM (
    'operational',
    'in_use',
    'under_repair',
    'written_off'
);

CREATE TABLE transport (
    transport_number integer NOT NULL PRIMARY KEY,
    id_vgk_location integer DEFAULT NULL,
    model varchar(255) NOT NULL,
    type varchar(255) NOT NULL,
    status transport_status_enum NOT NULL,
    manufacture_date date NOT NULL,
    mileage decimal NOT NULL,
    last_check_date date NOT NULL,

    CONSTRAINT manufacture_date_check CHECK (
        manufacture_date BETWEEN '2000-01-01' AND current_date
    ),

    CONSTRAINT mileage_check CHECK (
        mileage BETWEEN 0.00 AND 1000000.00
    ),

    CONSTRAINT last_check_date_check CHECK (
        last_check_date BETWEEN manufacture_date AND current_date
    ),

    CONSTRAINT transport_id_vgk_location_vgk_locations_id_vgk_location_foreign FOREIGN KEY (id_vgk_location) REFERENCES vgk_locations (id_vgk_location)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE equipment_usage_history (
    inventory_number integer NOT NULL,
    id_rescuer integer NOT NULL,
    issue_date date NOT NULL,
    return_date date NOT NULL,
    purpose varchar(255) NOT NULL,

    PRIMARY KEY (inventory_number, id_rescuer),

    CONSTRAINT issue_date_check CHECK (
        issue_date <= current_date
    ),

    CONSTRAINT return_date_check CHECK (
        return_date > issue_date AND return_date <= current_date
    ),

    CONSTRAINT purpose_not_empty_check CHECK (
        length(trim(purpose)) > 0
    ),

    CONSTRAINT equipment_usage_history_id_rescuer_vgk_rescuers_id_rescuer_foreign FOREIGN KEY (id_rescuer) REFERENCES vgk_rescuers (id_rescuer)
        ON UPDATE CASCADE
        ON DELETE CASCADE,

    CONSTRAINT equipment_usage_history_inventory_number_equipment_inventory_number_foreign FOREIGN KEY (inventory_number) REFERENCES equipment (inventory_number)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE transport_usage_history (
    transport_number integer NOT NULL,
    id_rescuer integer NOT NULL,
    departure_date date NOT NULL,
    return_date date NOT NULL,
    purpose varchar(255) NOT NULL,

    PRIMARY KEY (transport_number, id_rescuer),

    CONSTRAINT departure_date_check CHECK (
        departure_date <= current_date
    ),

    CONSTRAINT return_date_check CHECK (
        return_date > departure_date AND return_date <= current_date
    ),

    CONSTRAINT purpose_not_empty_check CHECK (
        length(trim(purpose)) > 0
    ),

    CONSTRAINT transport_usage_history_id_rescuer_vgk_rescuers_id_rescuer_foreign FOREIGN KEY (id_rescuer) REFERENCES vgk_rescuers (id_rescuer)
        ON UPDATE CASCADE
        ON DELETE CASCADE,

    CONSTRAINT transport_usage_history_transport_number_transport_transport_number_foreign FOREIGN KEY (transport_number) REFERENCES transport (transport_number)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TYPE service_status_enum AS ENUM (
    'in_queue',
    'under_repair',
    'repaired',
    'irreparable'
);

CREATE TABLE equipment_service_history (
    inventory_number integer NOT NULL,
    id_service_room integer NOT NULL,
    reason varchar(255) NOT NULL,
    serve_date date NOT NULL,
    status service_status_enum NOT NULL,

    PRIMARY KEY (
        inventory_number, id_service_room
    ),

    CONSTRAINT serve_date_check CHECK (
        serve_date <= current_date
    ),

    CONSTRAINT reason_not_empty_check CHECK (
        length(trim(reason)) > 0
    ),

    CONSTRAINT equipment_service_history_inventory_number_equipment_inventory_number_foreign FOREIGN KEY (inventory_number) REFERENCES equipment (inventory_number)
        ON UPDATE CASCADE
        ON DELETE CASCADE,

    CONSTRAINT equipment_service_history_id_service_room_vgk_service_room_id_service_room_foreign FOREIGN KEY (id_service_room) REFERENCES vgk_service_room (id_service_room)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);

CREATE TABLE transport_service_history (
    transport_number integer NOT NULL,
    id_service_room integer NOT NULL,
    reason varchar(255) NOT NULL,
    serve_date date NOT NULL,
    status service_status_enum NOT NULL,

    PRIMARY KEY (
        transport_number, id_service_room
    ),

    CONSTRAINT serve_date_check CHECK (
        serve_date <= current_date
    ),

    CONSTRAINT reason_not_empty_check CHECK (
        length(trim(reason)) > 0
    ),

    CONSTRAINT transport_service_history_id_service_room_vgk_service_room_id_service_room_foreign FOREIGN KEY (id_service_room) REFERENCES vgk_service_room (id_service_room)
        ON UPDATE CASCADE
        ON DELETE CASCADE,

    CONSTRAINT transport_service_history_transport_number_transport_transport_number_foreign FOREIGN KEY (transport_number) REFERENCES transport (transport_number)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);
