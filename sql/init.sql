CREATE TABLE equipment_types (
  type_name varchar(255) NOT NULL,
  equipment_standards_url varchar(255) NOT NULL,
  PRIMARY KEY (type_name)
);

CREATE TABLE objects (
  id_object integer NOT NULL,
  object_type varchar(255) NOT NULL,
  name varchar(255) NOT NULL,
  address varchar(255) NOT NULL,
  phone varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  director_full_name varchar(255) NOT NULL,
  PRIMARY KEY (id_object)
);

CREATE TABLE accidents (
  id_accident integer NOT NULL,
  id_object integer NOT NULL,
  accident_type varchar(255) NOT NULL,
  begin_date_time timestamp NOT NULL,
  status varchar(255) NOT NULL,
  description text NOT NULL,
  first_estimate text NOT NULL,
  cause text NOT NULL,
  PRIMARY KEY (id_accident),
  CONSTRAINT accidents_id_object_objects_id_object_foreign FOREIGN KEY (id_object) REFERENCES objects (id_object)
    ON UPDATE CASCADE
    ON DELETE RESTRICT
);

CREATE TABLE applications_for_admission (
  id_application integer NOT NULL,
  id_object integer NOT NULL,
  passport_number integer NOT NULL,
  first_name varchar(255) NOT NULL,
  last_name varchar(255) NOT NULL,
  surname varchar(255) DEFAULT NULL,
  issue_date date NOT NULL,
  phone varchar(255) NOT NULL,
  email varchar(255) DEFAULT NULL,
  status varchar(255) NOT NULL,
  birthday_date date NOT NULL,
  PRIMARY KEY (id_application),
  CONSTRAINT applications_for_admission_id_object_objects_id_object_foreign FOREIGN KEY (id_object) REFERENCES objects (id_object)
    ON UPDATE CASCADE
    ON DELETE RESTRICT
);

CREATE TABLE candidates_documents (
  document_type varchar(255) NOT NULL,
  id_application integer NOT NULL,
  document_url varchar(255) NOT NULL,
  valid_until date NOT NULL,
  PRIMARY KEY (document_type, id_application),
  CONSTRAINT candidates_documents_id_application_applications_for_admission_id_application_foreign FOREIGN KEY (id_application) REFERENCES applications_for_admission (id_application)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE candidates_medical_parameters (
  id_application integer NOT NULL,
  date date NOT NULL,
  expire_date date NOT NULL,
  health_group integer NOT NULL,
  height decimal NOT NULL,
  weight decimal NOT NULL,
  note text NOT NULL,
  PRIMARY KEY (id_application),
  CONSTRAINT candidates_medical_parameters_id_application_applications_for_admission_id_application_foreign FOREIGN KEY (id_application) REFERENCES applications_for_admission (id_application)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE vgk (
  id_vgk integer NOT NULL,
  id_object integer NOT NULL,
  status varchar(255) NOT NULL,
  formation_date date NOT NULL,
  PRIMARY KEY (id_vgk),
  CONSTRAINT vgk_id_object_objects_id_object_foreign FOREIGN KEY (id_object) REFERENCES objects (id_object)
    ON UPDATE CASCADE
    ON DELETE RESTRICT
);

CREATE TABLE positions (
  position_name varchar(255) NOT NULL,
  salary decimal(10, 2) NOT NULL,
  min_experience_years integer DEFAULT 0,
  responsibilities text,
  PRIMARY KEY (position_name)
);

CREATE TABLE vgk_rescuers (
  id_rescuer integer NOT NULL,
  id_vgk integer DEFAULT NULL,
  position varchar(255) DEFAULT NULL,
  first_name varchar(255) NOT NULL,
  second_name varchar(255) NOT NULL,
  surname varchar(255) NOT NULL,
  status varchar(255) NOT NULL,
  birth_date date NOT NULL,
  home_address varchar(255) NOT NULL,
  experience_years integer NOT NULL DEFAULT 0,
  PRIMARY KEY (id_rescuer),
  CONSTRAINT vgk_rescuers_id_vgk_vgk_id_vgk_foreign FOREIGN KEY (id_vgk) REFERENCES vgk (id_vgk)
    ON UPDATE CASCADE
    ON DELETE RESTRICT,
  CONSTRAINT position_foreign FOREIGN KEY (position) REFERENCES positions (position_name)
    ON UPDATE CASCADE
    ON DELETE RESTRICT,
  CONSTRAINT commander_must_be_in_vgk CHECK (
    NOT (position = 'commander' AND id_vgk IS NULL)
  )
);

CREATE TABLE vgk_rescuers_documents (
  document_type varchar(255) NOT NULL,
  id_rescuer integer NOT NULL,
  document_url varchar(255) NOT NULL,
  valid_until date NOT NULL,
  PRIMARY KEY (document_type, id_rescuer),
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
  CONSTRAINT vgk_shifts_id_vgk_vgk_id_vgk_foreign FOREIGN KEY (id_vgk) REFERENCES vgk (id_vgk)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
  CONSTRAINT vgk_shifts_id_vgk_location_vgk_locations_id_vgk_location_foreign FOREIGN KEY (id_vgk_location) REFERENCES vgk_locations (id_vgk_location)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE accidents_response_operations (
  id_operation integer NOT NULL,
  id_accident integer NOT NULL,
  id_leader integer NOT NULL,
  start_date_time timestamp NOT NULL,
  end_date_time timestamp DEFAULT NULL,
  status varchar(255) NOT NULL,
  PRIMARY KEY (id_operation),
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
  CONSTRAINT operations_participations_id_vgk_vgk_id_vgk_foreign FOREIGN KEY (id_vgk) REFERENCES vgk (id_vgk)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
  CONSTRAINT operations_participations_id_operation_accidents_response_operations_id_operation_foreign FOREIGN KEY (id_operation) REFERENCES accidents_response_operations (id_operation)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE operations_reports (
  id_report integer NOT NULL,
  id_operation integer NOT NULL,
  report_date_time timestamp NOT NULL,
  description text NOT NULL,
  PRIMARY KEY (id_report),
  CONSTRAINT operations_reports_id_participation_accidents_response_operations_id_operation_foreign FOREIGN KEY (id_operation) REFERENCES accidents_response_operations (id_operation)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE trainings (
  date date NOT NULL,
  id_object_location integer NOT NULL,
  id_instructor integer NOT NULL,
  topic varchar(255) NOT NULL,
  description text NOT NULL,
  PRIMARY KEY (date, id_object_location),
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
  notes text NOT NULL,
  PRIMARY KEY (date, id_rescuer),
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
  note text NOT NULL,
  PRIMARY KEY (date, id_rescuer),
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
  CONSTRAINT vgk_service_room_id_responsible_vgk_rescuers_id_rescuer_foreign FOREIGN KEY (id_responsible) REFERENCES vgk_rescuers (id_rescuer)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE vgk_locations (
  id_vgk_location integer NOT NULL,
  id_responsible integer DEFAULT NULL,
  address varchar(255) NOT NULL,
  status varchar(255) NOT NULL,
  PRIMARY KEY (id_vgk_location),
  CONSTRAINT vgk_locations_id_responsible_vgk_rescuers_id_rescuer_foreign FOREIGN KEY (id_responsible) REFERENCES vgk_rescuers (id_rescuer)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE equipment (
  inventory_number integer NOT NULL,
  id_vgk_location integer DEFAULT NULL,
  equipment_type varchar(255) NOT NULL,
  name varchar(255) NOT NULL,
  status varchar(255) NOT NULL,
  last_check_date date NOT NULL,
  PRIMARY KEY (inventory_number),
  CONSTRAINT equipment_id_vgk_location_vgk_locations_id_vgk_location_foreign FOREIGN KEY (id_vgk_location) REFERENCES vgk_locations (id_vgk_location)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
  CONSTRAINT equipment_equipment_type_equipment_types_type_name_foreign FOREIGN KEY (equipment_type) REFERENCES equipment_types (type_name)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE transport (
  transport_number integer NOT NULL,
  id_vgk_location integer DEFAULT NULL,
  model varchar(255) NOT NULL,
  type varchar(255) NOT NULL,
  status varchar(255) NOT NULL,
  manufacture_date date NOT NULL,
  mileage decimal NOT NULL,
  last_check_date date NOT NULL,
  PRIMARY KEY (transport_number),
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
  CONSTRAINT transport_usage_history_id_rescuer_vgk_rescuers_id_rescuer_foreign FOREIGN KEY (id_rescuer) REFERENCES vgk_rescuers (id_rescuer)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
  CONSTRAINT transport_usage_history_transport_number_transport_transport_number_foreign FOREIGN KEY (transport_number) REFERENCES transport (transport_number)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE equipment_service_history (
  inventory_number integer NOT NULL,
  id_service_room integer NOT NULL,
  reason varchar(255) NOT NULL,
  serve_date date NOT NULL,
  status varchar(255) NOT NULL,
  PRIMARY KEY (
    inventory_number, id_service_room
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
  status varchar(255) NOT NULL,
  PRIMARY KEY (
    transport_number, id_service_room
  ),
  CONSTRAINT transport_service_history_id_service_room_vgk_service_room_id_service_room_foreign FOREIGN KEY (id_service_room) REFERENCES vgk_service_room (id_service_room)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
  CONSTRAINT transport_service_history_transport_number_transport_transport_number_foreign FOREIGN KEY (transport_number) REFERENCES transport (transport_number)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);
