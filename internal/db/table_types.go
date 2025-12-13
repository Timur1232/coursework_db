package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/jackc/pgx/v5"
)

// ======================[types]====================== //

type Role string

const (
	Role_Guest     Role = "guest"
	Role_Candidate Role = "candidate"
	Role_Rescuer   Role = "rescuer"
	Role_Operator  Role = "operator"
	Role_Admin     Role = "admin"
)

type Users struct {
	IdUser   uint64 `db:"id_user"`
	Login    string `db:"login"`
	Password string `db:"password"`
	Role     Role   `db:"role"`
}

type EquipmentTypes struct {
	TypeName              string `db:"type_name"`
	EquipmentStandardsUrl string `db:"equipment_standards_url"`
}

type ObjectType string

const (
	ObjectType_Surface          ObjectType = "surface"
	ObjectType_Deep             ObjectType = "deep"
	ObjectType_Upland           ObjectType = "upland"
	ObjectType_UplandDeep       ObjectType = "upland_deep"
	ObjectType_UnderwaterMining ObjectType = "underwater_mining"
)

type Objects struct {
	IdObject         uint64     `db:"id_object"`
	ObjectType       ObjectType `db:"object_type"`
	Name             string     `db:"name"`
	Address          string     `db:"address"`
	Phone            string     `db:"phone"`
	Email            string     `db:"email"`
	DirectorFullName string     `db:"director_full_name"`
}

type DangerLevel string

const (
	DangerLevel_Low      DangerLevel = "low"
	DangerLevel_Medium   DangerLevel = "medium"
	DangerLevel_High     DangerLevel = "high"
	DangerLevel_Critical DangerLevel = "critical"
)

type AccidentTypes struct {
	AccidentName string      `db:"accident_name"`
	DangerLevel  DangerLevel `db:"danger_level"`
}

type AccidentStatus string

const (
	AccidentStatus_Reported      AccidentStatus = "reported"
	AccidentStatus_InProgress    AccidentStatus = "in_progress"
	AccidentStatus_Contained     AccidentStatus = "contained"
	AccidentStatus_Resolved      AccidentStatus = "resolved"
	AccidentStatus_Investigating AccidentStatus = "investigating"
)

type Accidents struct {
	IdAccident    uint64         `db:"id_accident"`
	IdObject      uint64         `db:"id_object"`
	AccidentType  string         `db:"accident_type"`
	BeginDateTime time.Time      `db:"begin_date_time"`
	Status        AccidentStatus `db:"status"`
	Description   string         `db:"description"`
	FirstEstimate string         `db:"first_estimate"`
	Cause         string         `db:"cause"`
}

type ApplicationsForAdmission struct {
	IdApplication  uint64         `db:"id_application"`
	IdObject       uint64         `db:"id_object"`
	PassportNumber string         `db:"passport_number"`
	FirstName      string         `db:"first_name"`
	LastName       string         `db:"last_name"`
	Surname        sql.NullString `db:"surname"`
	IssueDate      time.Time      `db:"issue_date"`
	Phone          string         `db:"phone"`
	Email          string         `db:"email"`
	Status         string         `db:"status"`
	BirthdayDate   time.Time      `db:"birthday_date"`
	HomeAddress    string         `db:"home_address"`
	IdUser         sql.NullInt64  `db:"id_user"`
}

type DocumentType string

const (
	DocumentType_MedicalCertificate  DocumentType = "medical_certificate"
	DocumentType_EducationDiploma    DocumentType = "education_diploma"
	DocumentType_TrainingCertificate DocumentType = "training_certificate"
	DocumentType_WorkBook            DocumentType = "work_book"
	DocumentType_MilitaryId          DocumentType = "military_id"
	DocumentType_PassportCopy        DocumentType = "passport_copy"
	DocumentType_Photo               DocumentType = "photo"
)

type CandidatesDocuments struct {
	DocumentType  DocumentType `db:"document_type"`
	IdApplication uint64       `db:"id_application"`
	DocumentUrl   string       `db:"document_url"`
	ValidUntil    time.Time    `db:"valid_until"`
}

type CandidatesMedicalParameters struct {
	IdApplication uint64    `db:"id_application"`
	Date          time.Time `db:"date"`
	ExpireDate    time.Time `db:"expire_date"`
	HealthGroup   string    `db:"health_group"`
	Height        float32   `db:"height"`
	Weight        float32   `db:"weight"`
	Note          string    `db:"note"`
}

type VgkStatus string

const (
	VgkStatus_OnDuty      VgkStatus = "on_duty"
	VgkStatus_Dismissed   VgkStatus = "dismissed"
	VgkStatus_OnDeparture VgkStatus = "on_departure"
	VgkStatus_OnShift     VgkStatus = "on_shift"
	VgkStatus_Inactive    VgkStatus = "inactive"
)

type Vgk struct {
	IdVgk         uint64    `db:"id_vgk"`
	IdObject      uint64    `db:"id_object"`
	Status        VgkStatus `db:"status"`
	FormationDate time.Time `db:"formation_date"`
}

type Positions struct {
	PositionName       string  `db:"position_name"`
	Salary             float32 `db:"salary"`
	MinExperienceYears uint    `db:"min_experience_years"`
	Responsibilities   string  `db:"responsibilities"`
}

type VgkRescuers struct {
	IdRescuer       uint64         `db:"id_rescuer"`
	IdVgk           sql.NullInt64  `db:"id_vgk"`
	Position        sql.NullString `db:"position"`
	FirstName       string         `db:"first_name"`
	SecondName      string         `db:"second_name"`
	Surname         sql.NullString `db:"surname"`
	Status          VgkStatus      `db:"status"`
	BirthDate       time.Time      `db:"birth_date"`
	HomeAddress     string         `db:"home_address"`
	ExperienceYears uint           `db:"experience_years"`
	IdUser          sql.NullInt64  `db:"id_user"`
}

type VgkRescuersDocuments struct {
	DocumentType string    `db:"document_type"`
	IdRescuer    uint64    `db:"id_rescuer"`
	DocumentUrl  string    `db:"document_url"`
	ValidUntil   time.Time `db:"valid_until"`
}

type VgkLocationStatus string

const (
	VgkLocationStatus_Operational    VgkLocationStatus = "operational"
	VgkLocationStatus_Malfunctioning VgkLocationStatus = "malfunctioning"
)

type VgkLocations struct {
	IdVgkLocation uint64            `db:"id_vgk_location"`
	IdResponsible sql.NullInt64     `db:"id_responsible"`
	Address       string            `db:"address"`
	Status        VgkLocationStatus `db:"status"`
}

type VgkShifts struct {
	ShiftStart    time.Time `db:"shift_start"`
	IdVgk         uint64    `db:"id_vgk"`
	IdVgkLocation uint64    `db:"id_vgk_location"`
	ShiftEnd      time.Time `db:"shift_end"`
}

type OperationStatus string

const (
	OperationStatus_Assessment        OperationStatus = "assessment"
	OperationStatus_Evacuation        OperationStatus = "evacuation"
	OperationStatus_Aftermath_cleanup OperationStatus = "aftermath_cleanup"
	OperationStatus_Safety_ensuring   OperationStatus = "safety_ensuring"
	OperationStatus_Completed         OperationStatus = "completed"
	OperationStatus_Failed            OperationStatus = "failed"
)

type AccidentsResponseOperations struct {
	IdOperation   uint64          `db:"id_operation"`
	IdAccident    uint64          `db:"id_accident"`
	StartDateTime time.Time       `db:"start_date_time"`
	EndDateTime   sql.NullTime    `db:"end_date_time"`
	Status        OperationStatus `db:"status"`
}

type OperationsParticipations struct {
	IdVgk        uint64 `db:"id_vgk"`
	IdOperation  uint64 `db:"id_operation"`
	AssignedTask string `db:"assigned_task"`
}

type OperationsReports struct {
	IdReport       uint64    `db:"id_report"`
	IdOperation    uint64    `db:"id_operation"`
	ReportDateTime time.Time `db:"report_date_time"`
	Description    string    `db:"description"`
}

type Trainings struct {
	Date             time.Time      `db:"date"`
	IdObjectLocation uint64         `db:"id_object_location"`
	IdInstructor     uint64         `db:"id_instructor"`
	Topic            string         `db:"topic"`
	Description      sql.NullString `db:"description"`
}

type TrainingsParticipations struct {
	Date             time.Time      `db:"date"`
	IdObjectLocation uint64         `db:"id_object_location"`
	IdRescuer        uint64         `db:"id_rescuer"`
	Notes            sql.NullString `db:"notes"`
}

type CertificationsPassings struct {
	Date      time.Time `db:"date"`
	IdRescuer uint64    `db:"id_rescuer"`
	Result    bool      `db:"result"`
	Topic     string    `db:"topic"`
}

type VgkRescuersMedicalParameters struct {
	Date        time.Time      `db:"date"`
	IdRescuer   uint64         `db:"id_rescuer"`
	ExpireDate  time.Time      `db:"expire_date"`
	HealthGroup uint           `db:"health_group"`
	Height      float32        `db:"height"`
	Weight      float32        `db:"weight"`
	Conclusion  string         `db:"conclusion"`
	Note        sql.NullString `db:"note"`
}

type VgkServiceRoom struct {
	IdServiceRoom uint64        `db:"id_service_room"`
	IdResponsible sql.NullInt64 `db:"id_responsible"`
	Purpose       string        `db:"purpose"`
	Address       string        `db:"address"`
}

type EquipmentStatus string

const (
	EquipmentStatus_Operational        EquipmentStatus = "operational"
	EquipmentStatus_InUse              EquipmentStatus = "in_use"
	EquipmentStatus_NeedsRepairService EquipmentStatus = "needs_repair_service"
	EquipmentStatus_UnderRepair        EquipmentStatus = "under_repair"
	EquipmentStatus_WrittenOff         EquipmentStatus = "written_off"
)

type Equipment struct {
	InventoryNumber uint64          `db:"inventory_number"`
	IdVgkLocation   sql.NullInt64   `db:"id_vgk_location"`
	EquipmentType   string          `db:"equipment_type"`
	Name            string          `db:"name"`
	Status          EquipmentStatus `db:"status"`
	LastCheckDate   time.Time       `db:"last_check_date"`
}

type Transport struct {
	TransportNumber uint64          `db:"transport_number"`
	IdVgkLocation   sql.NullInt64   `db:"id_vgk_location"`
	Model           string          `db:"model"`
	Type            string          `db:"type"`
	Status          EquipmentStatus `db:"status"`
	ManufactureDate time.Time       `db:"manufacture_date"`
	Mileage         float32         `db:"mileage"`
	LastCheckDate   time.Time       `db:"last_check_date"`
}

type EquipmentUsageHistory struct {
	InventoryNumber uint64    `db:"inventory_number"`
	IdRescuer       uint64    `db:"id_rescuer"`
	IssueDate       time.Time `db:"issue_date"`
	ReturnDate      time.Time `db:"return_date"`
	Purpose         string    `db:"purpose"`
}

type TransportUsageHistory struct {
	TransportNumber uint64    `db:"transport_number"`
	IdRescuer       uint64    `db:"id_rescuer"`
	DepartureDate   time.Time `db:"departure_date"`
	ReturnDate      time.Time `db:"return_date"`
	Purpose         string    `db:"purpose"`
}

type ServiceStatus string

const (
	ServeceStatus_InQueue     ServiceStatus = "in_queue"
	ServeceStatus_UnderRepair ServiceStatus = "under_repair"
	ServeceStatus_Repaired    ServiceStatus = "repaired"
	ServeceStatus_Irreparable ServiceStatus = "irreparable"
)

type EquipmentServiceHistory struct {
	InventoryNumber uint64        `db:"inventory_number"`
	IdServiceRoom   uint64        `db:"id_service_room"`
	Reason          string        `db:"reason"`
	ServeDate       time.Time     `db:"serve_date"`
	Status          ServiceStatus `db:"status"`
}

type TransportServiceHistory struct {
	TransportNumber uint64        `db:"transport_number"`
	IdServiceRoom   uint64        `db:"id_service_room"`
	Reason          string        `db:"reason"`
	ServeDate       time.Time     `db:"serve_date"`
	Status          ServiceStatus `db:"status"`
}

// ======================[funcs]====================== //

func Query[T any](conn *pgx.Conn, ctx context.Context, sql string, args ...any) (any, error) {
	rows, err := conn.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result, err := pgx.CollectRows(rows, pgx.RowToStructByName[T])
	if err != nil {
		return nil, err
	}

	return result, nil
}

/*func QueryRowEquipmentTypes(conn *pgx.Conn, ctx context.Context, sql string, args ...any) (*EquipmentTypes, error) {
	row := conn.QueryRow(ctx, sql, args...)

	var result EquipmentTypes
	err := row.Scan(&result.TypeName, &result.EquipmentStandardsUrl)
	if err != nil {
		return nil, err
	}

	return &result, nil
}*/

// EquipmentTypes
// Objects
// AccidentTypes
// Accidents
// ApplicationsForAdmission
// CandidatesDocuments
// CandidatesMedicalParameters
// Vgk
// Positions
// VgkRescuers
// VgkRescuersDocuments
// VgkLocations
// VgkShifts
// AccidentsResponseOperations
// OperationsParticipations
// OperationsReports
// Trainings
// TrainingsParticipations
// CertificationsPassings
// VgkRescuersMedicalParameters
// VgkServiceRoom
// Equipment
// Transport
// EquipmentUsageHistory
// TransportUsageHistory
// EquipmentServiceHistory
// TransportServiceHistory
