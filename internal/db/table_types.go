package db

import "time"

type EquipmentTypes struct {
	TypeName              string
	EquipmentStandardsUrl string
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
	IdObject         uint64
	ObjectType       ObjectType
	Name             string
	Address          string
	Phone            string
	Email            string
	DirectorFullName string
}

type DangerLevel string

const (
	DangerLevel_Low      DangerLevel = "low"
	DangerLevel_Medium   DangerLevel = "medium"
	DangerLevel_High     DangerLevel = "high"
	DangerLevel_Critical DangerLevel = "critical"
)

type AccidentTypes struct {
	AccidentName string
	DangerLevel  DangerLevel
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
	IdAccident    uint64
	IdObject      uint64
	AccidentType  string
	BeginDateTime time.Time
	Status        AccidentStatus
	Description   string
	FirstEstimate string
	Cause         string
}

type ApplicationsForAdmission struct {
	IdApplication  uint64
	IdObject       uint64
	PassportNumber string
	FirstName      string
	LastName       string
	Surname        string
	IssueDate      time.Time
	Phone          string
	Email          string
	Status         string
	BirthdayDate   time.Time
	HomeAddress    string
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
	DocumentType  DocumentType
	IdApplication uint64
	DocumentUrl   string
	ValidUntil    time.Time
}

type CandidatesMedicalParameters struct {
	IdApplication uint64
	Date          time.Time
	ExpireDate    time.Time
	HealthGroup   string
	Height        float32
	Weight        float32
	Note          string
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
	IdVgk         uint64
	IdObject      uint64
	Status        VgkStatus
	FormationDate time.Time
}

type Positions struct {
	PositionName       string
	Salary             float32
	MinExperienceYears uint
	Responsibilities   string
}

type VgkRescuers struct {
	IdRescuer       uint64
	IdVgk           uint64
	Position        string
	FirstName       string
	SecondName      string
	Surname         string
	Status          VgkStatus
	BirthDate       time.Time
	HomeAddress     string
	ExperienceYears uint
}

type VgkRescuersDocuments struct {
	DocumentType string
	IdRescuer    uint64
	DocumentUrl  string
	ValidUntil   time.Time
}

type VgkLocationStatus string

const (
	VgkLocationStatus_Operational    VgkLocationStatus = "operational"
	VgkLocationStatus_Malfunctioning VgkLocationStatus = "malfunctioning"
)

type VgkLocations struct {
	IdVgkLocation uint64
	IdResponsible uint64
	Address       string
	Status        VgkLocationStatus
}

type VgkShifts struct {
	ShiftStart    time.Time
	IdVgk         uint64
	IdVgkLocation uint64
	ShiftEnd      time.Time
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
	IdOperation   uint64
	IdAccident    uint64
	StartDateTime time.Time
	EndDateTime   time.Time
	Status        OperationStatus
}

type OperationsParticipations struct {
	IdVgk        uint64
	IdOperation  uint64
	AssignedTask string
}

type OperationsReports struct {
	IdReport       uint64
	IdOperation    uint64
	ReportDateTime time.Time
	Description    string
}

type Trainings struct {
	Date             time.Time
	IdObjectLocation uint64
	IdInstructor     uint64
	Topic            string
	Description      string
}

type TrainingsParticipations struct {
	Date             time.Time
	IdObjectLocation uint64
	IdRescuer        uint64
	Notes            string
}

type CertificationsPassings struct {
	Date      time.Time
	IdRescuer uint64
	Result    bool
	Topic     string
}

type VgkRescuersMedicalParameters struct {
	Date        time.Time
	IdRescuer   uint64
	ExpireDate  time.Time
	HealthGroup uint
	Height      float32
	Weight      float32
	Conclusion  string
	Note        string
}

type VgkServiceRoom struct {
	IdServiceRoom uint64
	IdResponsible uint64
	Purpose       string
	Address       string
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
	InventoryNumber uint64
	IdVgkLocation   uint64
	EquipmentType   string
	Name            string
	Status          EquipmentStatus
	LastCheckDate   time.Time
}

type Transport struct {
	TransportNumber uint64
	IdVgkLocation   uint64
	Model           string
	Type            string
	Status          EquipmentStatus
	ManufactureDate time.Time
	Mileage         float32
	LastCheckDate   time.Time
}

type EquipmentUsageHistory struct {
	InventoryNumber uint64
	IdRescuer       uint64
	IssueDate       time.Time
	ReturnDate      time.Time
	Purpose         string
}

type TransportUsageHistory struct {
	TransportNumber uint64
	IdRescuer       uint64
	DepartureDate   time.Time
	ReturnDate      time.Time
	Purpose         string
}

type ServiceStatus string

const (
	ServeceStatus_InQueue     ServiceStatus = "in_queue"
	ServeceStatus_UnderRepair ServiceStatus = "under_repair"
	ServeceStatus_Repaired    ServiceStatus = "repaired"
	ServeceStatus_Irreparable ServiceStatus = "irreparable"
)

type EquipmentServiceHistory struct {
	InventoryNumber uint64
	IdServiceRoom   uint64
	Reason          string
	ServeDate       time.Time
	Status          ServiceStatus
}

type TransportServiceHistory struct {
	TransportNumber uint64
	IdServiceRoom   uint64
	Reason          string
	ServeDate       time.Time
	Status          ServiceStatus
}
