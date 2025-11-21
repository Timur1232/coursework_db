CREATE TABLE certification (
  id_certification integer NOT NULL,
  date date NOT NULL,
  topic varchar(255) NOT NULL,
  PRIMARY KEY (id_certification)
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

CREATE TABLE vgk_service_rooms (
  id_service_room integer NOT NULL,
  purpose varchar(255) NOT NULL,
  address varchar(255) NOT NULL,
  PRIMARY KEY (id_service_room)
);

CREATE TABLE equipment_types (
  type_name varchar(255) NOT NULL,
  PRIMARY KEY (type_name)
);

CREATE TABLE accidents (
  id_accident integer NOT NULL,
  id_object integer NOT NULL,
  accident_type varchar(255) NOT NULL,
  begin_date_time date NOT NULL,
  status varchar(255) NOT NULL,
  description text NOT NULL,
  first_estimate text NOT NULL,
  cause text NOT NULL,
  PRIMARY KEY (id_accident),
  CONSTRAINT accidents_id_object_objects_id_object_foreign FOREIGN KEY (id_object) REFERENCES objects(id_object)
    ON UPDATE CASCADE
    ON DELETE CASCADE
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
  PRIMARY KEY (id_application),
  CONSTRAINT applications_for_admission_id_object_objects_id_object_foreign FOREIGN KEY (id_object) REFERENCES objects(id_object)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE equipment_standards (
  id_norm integer NOT NULL,
  equipment_type varchar(255) NOT NULL,
  document_url varchar(255) NOT NULL,
  PRIMARY KEY (id_norm),
  CONSTRAINT equipment_standards_equipment_type_equipment_types_type_name_foreign FOREIGN KEY (equipment_type) REFERENCES equipment_types (type_name)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE candidate_documents (
  document_type varchar(255) NOT NULL,
  id_application integer NOT NULL,
  document_url varchar(255) NOT NULL,
  date date NOT NULL,
  valid_until date NOT NULL,
  PRIMARY KEY (document_type, id_application),
  CONSTRAINT candidate_documents_id_application_applications_for_admission_id_application_foreign FOREIGN KEY (id_application) REFERENCES applications_for_admission (id_application)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE candidates (
  id_application integer NOT NULL,
  id_object integer NOT NULL,
  first_name varchar(255) NOT NULL,
  last_name varchar(255) NOT NULL,
  surname varchar(255) NOT NULL,
  phone varchar(255) NOT NULL,
  email varchar(255) DEFAULT NULL,
  status varchar(255) NOT NULL,
  birth_date date NOT NULL,
  PRIMARY KEY (id_application),
  CONSTRAINT candidates_id_object_objects_id_object_foreign FOREIGN KEY (id_object) REFERENCES objects(id_object)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
  CONSTRAINT candidates_id_application_applications_for_admission_id_application_foreign FOREIGN KEY (id_application) REFERENCES applications_for_admission (id_application)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE candidates_medical_parameters (
  id_candidate integer NOT NULL,
  date date NOT NULL,
  expire_date date NOT NULL,
  health_group integer NOT NULL,
  height decimal NOT NULL,
  weight decimal NOT NULL,
  note text NOT NULL,
  PRIMARY KEY (id_candidate),
  CONSTRAINT candidates_medical_parameters_id_candidate_candidates_id_candidate_foreign FOREIGN KEY (id_candidate) REFERENCES candidates (id_application)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE vgks (
  id_vgk integer NOT NULL,
  id_commander integer DEFAULT NULL,
  id_object integer NOT NULL,
  status varchar(255) NOT NULL,
  formation_date date NOT NULL,
  PRIMARY KEY (id_vgk),
  CONSTRAINT vgks_id_object_objects_id_object_foreign FOREIGN KEY (id_object) REFERENCES objects(id_object)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
  CONSTRAINT vgks_id_commander_rescuers_id_rescuer_foreign FOREIGN KEY (id_commander) REFERENCES rescuers (id_rescuer)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE enrollment_orders (
  id_rescuer integer NOT NULL,
  id_vgk integer NOT NULL,
  position varchar(255) NOT NULL,
  date date NOT NULL,
  PRIMARY KEY (id_rescuer),
  CONSTRAINT enrollment_orders_id_vgk_vgks_id_vgk_foreign FOREIGN KEY (id_vgk) REFERENCES vgks (id_vgk)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
  CONSTRAINT enrollment_orders_id_rescuer_rescuers_id_rescuer_foreign FOREIGN KEY (id_rescuer) REFERENCES rescuers (id_rescuer)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE rescuers (
  id_rescuer integer NOT NULL,
  id_vgk integer NOT NULL,
  id_enrollment_order integer NOT NULL,
  position varchar(255) NOT NULL,
  first_name varchar(255) NOT NULL,
  second_name varchar(255) NOT NULL,
  surname varchar(255) NOT NULL,
  status varchar(255) NOT NULL,
  birth_date date NOT NULL,
  home_address varchar(255) NOT NULL,
  PRIMARY KEY (id_rescuer),
  CONSTRAINT rescuers_id_vgk_vgks_id_vgk_foreign FOREIGN KEY (id_vgk) REFERENCES vgks (id_vgk)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE rescuers_documents (
  document_type varchar(255) NOT NULL,
  id_rescuer integer NOT NULL,
  document_url varchar(255) NOT NULL,
  date date NOT NULL,
  valid_until date NOT NULL,
  PRIMARY KEY (document_type, id_rescuer),
  CONSTRAINT rescuers_documents_id_rescuer_rescuers_id_rescuer_foreign FOREIGN KEY (id_rescuer) REFERENCES rescuers (id_rescuer)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE shifts (
  shift_start date NOT NULL,
  id_vgk integer NOT NULL,
  id_vgk_location integer NOT NULL,
  shift_end date NOT NULL,
  PRIMARY KEY (
    shift_start, id_vgk, id_vgk_location
  ),
  CONSTRAINT shifts_id_vgk_vgks_id_vgk_foreign FOREIGN KEY (id_vgk) REFERENCES vgks (id_vgk)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
  CONSTRAINT shifts_id_vgk_location_vgk_locations_id_vgk_location_foreign FOREIGN KEY (id_vgk_location) REFERENCES vgk_locations (id_vgk_location)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE accident_response_operations (
  id_operation integer NOT NULL,
  id_accident integer NOT NULL,
  id_leader integer NOT NULL,
  start_date_time date NOT NULL,
  end_date_time date DEFAULT NULL,
  status varchar(255) NOT NULL,
  PRIMARY KEY (id_operation),
  CONSTRAINT accident_response_operations_id_accident_accidents_id_accident_foreign FOREIGN KEY (id_accident) REFERENCES accidents (id_accident)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
  CONSTRAINT accident_response_operations_id_leader_rescuers_id_rescuer_foreign FOREIGN KEY (id_leader) REFERENCES rescuers (id_rescuer)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE operation_participation (
  id_vgk integer NOT NULL,
  id_accident integer NOT NULL,
  status varchar(255) NOT NULL,
  assigned_tast varchar(255) NOT NULL,
  PRIMARY KEY (id_vgk, id_accident),
  CONSTRAINT operation_participation_id_vgk_vgks_id_vgk_foreign FOREIGN KEY (id_vgk) REFERENCES vgks (id_vgk)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
  CONSTRAINT operation_participation_id_accident_accident_response_operations_id_operation_foreign FOREIGN KEY (id_accident) REFERENCES accident_response_operations (id_operation)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE operation_reports (
  id_report integer NOT NULL,
  id_participation integer NOT NULL,
  report_date_time date NOT NULL,
  description text NOT NULL,
  PRIMARY KEY (id_report),
  CONSTRAINT operation_reports_id_participation_accident_response_operations_id_operation_foreign FOREIGN KEY (id_participation) REFERENCES accident_response_operations (id_operation)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE trainings (
  id_training integer NOT NULL,
  id_object_location integer NOT NULL,
  id_instructor integer NOT NULL,
  date date NOT NULL,
  topic varchar(255) NOT NULL,
  description text NOT NULL,
  PRIMARY KEY (id_training),
  CONSTRAINT Тренировка_id_instructor_rescuers_id_rescuer_foreign FOREIGN KEY (id_instructor) REFERENCES rescuers (id_rescuer)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
  CONSTRAINT trainings_id_object_location_objects_id_object_foreign FOREIGN KEY (id_object_location) REFERENCES objects(id_object)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE training_participation (
  id_training integer NOT NULL,
  id_rescuer integer NOT NULL,
  notes text NOT NULL,
  PRIMARY KEY (id_training, id_rescuer),
  CONSTRAINT training_participation_id_training_Тренировка_id_training_foreign FOREIGN KEY (id_training) REFERENCES trainings (id_training)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
  CONSTRAINT training_participation_id_rescuer_rescuers_id_rescuer_foreign FOREIGN KEY (id_rescuer) REFERENCES rescuers (id_rescuer)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE certification_passing (
  id_certification integer NOT NULL,
  id_rescuer integer NOT NULL,
  date date NOT NULL,
  result boolean NOT NULL,
  PRIMARY KEY (id_certification, id_rescuer),
  CONSTRAINT certification_passing_id_rescuer_rescuers_id_rescuer_foreign FOREIGN KEY (id_rescuer) REFERENCES rescuers (id_rescuer)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
  CONSTRAINT certification_passing_id_certification_certification_id_certification_foreign FOREIGN KEY (id_certification) REFERENCES certification (id_certification)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE rescuers_medical_parameters (
  id_medical_record integer NOT NULL,
  id_rescuer integer NOT NULL,
  date date NOT NULL,
  expire_date date NOT NULL,
  health_group integer NOT NULL,
  height decimal NOT NULL,
  weight decimal NOT NULL,
  conclusion varchar(255) NOT NULL,
  note text NOT NULL,
  PRIMARY KEY (id_medical_record, id_rescuer),
  CONSTRAINT rescuers_medical_parameters_id_rescuer_rescuers_id_rescuer_foreign FOREIGN KEY (id_rescuer) REFERENCES rescuers (id_rescuer)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE vgk_locations (
  id_vgk_location integer NOT NULL,
  id_responsible integer DEFAULT NULL,
  address varchar(255) NOT NULL,
  status varchar(255) NOT NULL,
  PRIMARY KEY (id_vgk_location),
  CONSTRAINT vgk_locations_id_responsible_rescuers_id_rescuer_foreign FOREIGN KEY (id_responsible) REFERENCES rescuers (id_rescuer)
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
  next_check_date date NOT NULL,
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
  PRIMARY KEY (transport_number),
  CONSTRAINT transport_id_vgk_location_vgk_locations_id_vgk_location_foreign FOREIGN KEY (id_vgk_location) REFERENCES vgk_locations (id_vgk_location)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE equipment_usage_history (
  id_equipment integer NOT NULL,
  id_rescuer integer NOT NULL,
  issue_date date NOT NULL,
  return_date date NOT NULL,
  purpose varchar(255) NOT NULL,
  notes text NOT NULL,
  PRIMARY KEY (id_equipment, id_rescuer),
  CONSTRAINT equipment_usage_history_id_equipment_equipment_inventory_number_foreign FOREIGN KEY (id_equipment) REFERENCES equipment (inventory_number)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
  CONSTRAINT equipment_usage_history_id_rescuer_rescuers_id_rescuer_foreign FOREIGN KEY (id_rescuer) REFERENCES rescuers (id_rescuer)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE transport_usage_history (
  id_vehicle integer NOT NULL,
  id_rescuer integer NOT NULL,
  departure_date date NOT NULL,
  return_date date NOT NULL,
  purpose varchar(255) NOT NULL,
  notes text NOT NULL,
  PRIMARY KEY (id_vehicle, id_rescuer),
  CONSTRAINT transport_usage_history_id_vehicle_transport_transport_number_foreign FOREIGN KEY (id_vehicle) REFERENCES transport (transport_number)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
  CONSTRAINT transport_usage_history_id_rescuer_rescuers_id_rescuer_foreign FOREIGN KEY (id_rescuer) REFERENCES rescuers (id_rescuer)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE equipment_service_history (
  id_equipment integer NOT NULL,
  id_service_room integer NOT NULL,
  reason varchar(255) NOT NULL,
  serve_date date NOT NULL,
  status varchar(255) NOT NULL,
  notes text DEFAULT NULL,
  PRIMARY KEY (id_equipment, id_service_room),
  CONSTRAINT equipment_service_history_id_equipment_equipment_inventory_number_foreign FOREIGN KEY (id_equipment) REFERENCES equipment (inventory_number)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
  CONSTRAINT equipment_service_history_id_service_room_vgk_service_rooms_id_service_room_foreign FOREIGN KEY (id_service_room) REFERENCES vgk_service_rooms (id_service_room)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE transport_service_history (
  id_vihicle integer NOT NULL,
  id_service_room integer NOT NULL,
  reason varchar(255) NOT NULL,
  serve_date date NOT NULL,
  status varchar(255) NOT NULL,
  notes text DEFAULT NULL,
  PRIMARY KEY (id_vihicle, id_service_room),
  CONSTRAINT transport_service_history_id_vihicle_transport_transport_number_foreign FOREIGN KEY (id_vihicle) REFERENCES transport (transport_number)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
  CONSTRAINT transport_service_history_id_service_room_vgk_service_rooms_id_service_room_foreign FOREIGN KEY (id_service_room) REFERENCES vgk_service_rooms (id_service_room)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);
