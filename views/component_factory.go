package views

import (
	"github.com/Timur1232/coursework_db/internal/db"
	"github.com/a-h/templ"
)

// Интерфейс фабрики компонентов
type TableComponentFactory interface {
    CreateListComponent(data any) templ.Component
    CreateSortComponent() templ.Component
    GetTableName() string
}

// ====================
type EquipmentTypesTableFactory struct{}

func (f *EquipmentTypesTableFactory) CreateListComponent(data any) templ.Component {
    equipmentTypes := data.([]db.EquipmentTypes)
    return TableListEquipmentTypes(equipmentTypes)
}

func (f *EquipmentTypesTableFactory) CreateSortComponent() templ.Component {
    return SortEquipmentTypes()
}

func (f *EquipmentTypesTableFactory) GetTableName() string {
    return "equipment_types"
}

// ====================
type ObjectsTableFactory struct{}

func (f *ObjectsTableFactory) CreateListComponent(data any) templ.Component {
    objects := data.([]db.Objects)
    return TableListObjects(objects)
}

func (f *ObjectsTableFactory) CreateSortComponent() templ.Component {
    return SortObjects()
}

func (f *ObjectsTableFactory) GetTableName() string {
    return "objects"
}

// ====================
type AccidentTypesTableFactory struct{}

func (f *AccidentTypesTableFactory) CreateListComponent(data any) templ.Component {
    accidentTypes := data.([]db.AccidentTypes)
    return TableListAccidentTypes(accidentTypes)
}

func (f *AccidentTypesTableFactory) CreateSortComponent() templ.Component {
    return SortAccidentTypes()
}

func (f *AccidentTypesTableFactory) GetTableName() string {
    return "accident_types"
}

// ====================
type AccidentsTableFactory struct{}

func (f *AccidentsTableFactory) CreateListComponent(data any) templ.Component {
    accidents := data.([]db.Accidents)
    return TableListAccidents(accidents)
}

func (f *AccidentsTableFactory) CreateSortComponent() templ.Component {
    return SortAccidents()
}

func (f *AccidentsTableFactory) GetTableName() string {
    return "accidents"
}

// ====================
type ApplicationsForAdmissionTableFactory struct{}

func (f *ApplicationsForAdmissionTableFactory) CreateListComponent(data any) templ.Component {
    applications := data.([]db.ApplicationsForAdmission)
    return TableListApplicationsForAdmission(applications)
}

func (f *ApplicationsForAdmissionTableFactory) CreateSortComponent() templ.Component {
    return SortApplicationsForAdmission()
}

func (f *ApplicationsForAdmissionTableFactory) GetTableName() string {
    return "applications_for_admission"
}

// ====================
type CandidatesDocumentsTableFactory struct{}

func (f *CandidatesDocumentsTableFactory) CreateListComponent(data any) templ.Component {
    documents := data.([]db.CandidatesDocuments)
    return TableListCandidatesDocuments(documents)
}

func (f *CandidatesDocumentsTableFactory) CreateSortComponent() templ.Component {
    return SortCandidatesDocuments()
}

func (f *CandidatesDocumentsTableFactory) GetTableName() string {
    return "candidates_documents"
}

// ====================
type CandidatesMedicalParametersTableFactory struct{}

func (f *CandidatesMedicalParametersTableFactory) CreateListComponent(data any) templ.Component {
    medicalParams := data.([]db.CandidatesMedicalParameters)
    return TableListCandidatesMedicalParameters(medicalParams)
}

func (f *CandidatesMedicalParametersTableFactory) CreateSortComponent() templ.Component {
    return SortCandidatesMedicalParameters()
}

func (f *CandidatesMedicalParametersTableFactory) GetTableName() string {
    return "candidates_medical_parameters"
}

// ====================
type VgkTableFactory struct{}

func (f *VgkTableFactory) CreateListComponent(data any) templ.Component {
    vgk := data.([]db.Vgk)
    return TableListVgk(vgk)
}

func (f *VgkTableFactory) CreateSortComponent() templ.Component {
    return SortVgk()
}

func (f *VgkTableFactory) GetTableName() string {
    return "vgk"
}

// ====================
type PositionsTableFactory struct{}

func (f *PositionsTableFactory) CreateListComponent(data any) templ.Component {
    positions := data.([]db.Positions)
    return TableListPositions(positions)
}

func (f *PositionsTableFactory) CreateSortComponent() templ.Component {
    return SortPositions()
}

func (f *PositionsTableFactory) GetTableName() string {
    return "positions"
}

// ====================
type VgkRescuersTableFactory struct{}

func (f *VgkRescuersTableFactory) CreateListComponent(data any) templ.Component {
    rescuers := data.([]db.VgkRescuers)
    return TableListVgkRescuers(rescuers)
}

func (f *VgkRescuersTableFactory) CreateSortComponent() templ.Component {
    return SortVgkRescuers()
}

func (f *VgkRescuersTableFactory) GetTableName() string {
    return "vgk_rescuers"
}

// ====================
type VgkRescuersDocumentsTableFactory struct{}

func (f *VgkRescuersDocumentsTableFactory) CreateListComponent(data any) templ.Component {
    documents := data.([]db.VgkRescuersDocuments)
    return TableListVgkRescuersDocuments(documents)
}

func (f *VgkRescuersDocumentsTableFactory) CreateSortComponent() templ.Component {
    return SortVgkRescuersDocuments()
}

func (f *VgkRescuersDocumentsTableFactory) GetTableName() string {
    return "vgk_rescuers_documents"
}

// ====================
type VgkLocationsTableFactory struct{}

func (f *VgkLocationsTableFactory) CreateListComponent(data any) templ.Component {
    locations := data.([]db.VgkLocations)
    return TableListVgkLocations(locations)
}

func (f *VgkLocationsTableFactory) CreateSortComponent() templ.Component {
    return SortVgkLocations()
}

func (f *VgkLocationsTableFactory) GetTableName() string {
    return "vgk_locations"
}

// ====================
type VgkShiftsTableFactory struct{}

func (f *VgkShiftsTableFactory) CreateListComponent(data any) templ.Component {
    shifts := data.([]db.VgkShifts)
    return TableListVgkShifts(shifts)
}

func (f *VgkShiftsTableFactory) CreateSortComponent() templ.Component {
    return SortVgkShifts()
}

func (f *VgkShiftsTableFactory) GetTableName() string {
    return "vgk_shifts"
}

// ====================
type AccidentsResponseOperationsTableFactory struct{}

func (f *AccidentsResponseOperationsTableFactory) CreateListComponent(data any) templ.Component {
    operations := data.([]db.AccidentsResponseOperations)
    return TableListAccidentsResponseOperations(operations)
}

func (f *AccidentsResponseOperationsTableFactory) CreateSortComponent() templ.Component {
    return SortAccidentsResponseOperations()
}

func (f *AccidentsResponseOperationsTableFactory) GetTableName() string {
    return "accidents_response_operations"
}

// ====================
type OperationsParticipationsTableFactory struct{}

func (f *OperationsParticipationsTableFactory) CreateListComponent(data any) templ.Component {
    participations := data.([]db.OperationsParticipations)
    return TableListOperationsParticipations(participations)
}

func (f *OperationsParticipationsTableFactory) CreateSortComponent() templ.Component {
    return SortOperationsParticipations()
}

func (f *OperationsParticipationsTableFactory) GetTableName() string {
    return "operations_participations"
}

// ====================
type OperationsReportsTableFactory struct{}

func (f *OperationsReportsTableFactory) CreateListComponent(data any) templ.Component {
    reports := data.([]db.OperationsReports)
    return TableListOperationsReports(reports)
}

func (f *OperationsReportsTableFactory) CreateSortComponent() templ.Component {
    return SortOperationsReports()
}

func (f *OperationsReportsTableFactory) GetTableName() string {
    return "operations_reports"
}

// ====================
type TrainingsTableFactory struct{}

func (f *TrainingsTableFactory) CreateListComponent(data any) templ.Component {
    trainings := data.([]db.Trainings)
    return TableListTrainings(trainings)
}

func (f *TrainingsTableFactory) CreateSortComponent() templ.Component {
    return SortTrainings()
}

func (f *TrainingsTableFactory) GetTableName() string {
    return "trainings"
}

// ====================
type TrainingsParticipationsTableFactory struct{}

func (f *TrainingsParticipationsTableFactory) CreateListComponent(data any) templ.Component {
    participations := data.([]db.TrainingsParticipations)
    return TableListTrainingsParticipations(participations)
}

func (f *TrainingsParticipationsTableFactory) CreateSortComponent() templ.Component {
    return SortTrainingsParticipations()
}

func (f *TrainingsParticipationsTableFactory) GetTableName() string {
    return "trainings_participations"
}

// ====================
type CertificationsPassingsTableFactory struct{}

func (f *CertificationsPassingsTableFactory) CreateListComponent(data any) templ.Component {
    certifications := data.([]db.CertificationsPassings)
    return TableListCertificationsPassings(certifications)
}

func (f *CertificationsPassingsTableFactory) CreateSortComponent() templ.Component {
    return SortCertificationsPassings()
}

func (f *CertificationsPassingsTableFactory) GetTableName() string {
    return "certifications_passings"
}

// ====================
type VgkRescuersMedicalParametersTableFactory struct{}

func (f *VgkRescuersMedicalParametersTableFactory) CreateListComponent(data any) templ.Component {
    medicalParams := data.([]db.VgkRescuersMedicalParameters)
    return TableListVgkRescuersMedicalParameters(medicalParams)
}

func (f *VgkRescuersMedicalParametersTableFactory) CreateSortComponent() templ.Component {
    return SortVgkRescuersMedicalParameters()
}

func (f *VgkRescuersMedicalParametersTableFactory) GetTableName() string {
    return "vgk_rescuers_medical_parameters"
}

// ====================
type VgkServiceRoomTableFactory struct{}

func (f *VgkServiceRoomTableFactory) CreateListComponent(data any) templ.Component {
    serviceRooms := data.([]db.VgkServiceRoom)
    return TableListVgkServiceRoom(serviceRooms)
}

func (f *VgkServiceRoomTableFactory) CreateSortComponent() templ.Component {
    return SortVgkServiceRoom()
}

func (f *VgkServiceRoomTableFactory) GetTableName() string {
    return "vgk_service_room"
}

// ====================
type EquipmentTableFactory struct{}

func (f *EquipmentTableFactory) CreateListComponent(data any) templ.Component {
    equipment := data.([]db.Equipment)
    return TableListEquipment(equipment)
}

func (f *EquipmentTableFactory) CreateSortComponent() templ.Component {
    return SortEquipment()
}

func (f *EquipmentTableFactory) GetTableName() string {
    return "equipment"
}

// ====================
type TransportTableFactory struct{}

func (f *TransportTableFactory) CreateListComponent(data any) templ.Component {
    transport := data.([]db.Transport)
    return TableListTransport(transport)
}

func (f *TransportTableFactory) CreateSortComponent() templ.Component {
    return SortTransport()
}

func (f *TransportTableFactory) GetTableName() string {
    return "transport"
}

// ====================
type EquipmentUsageHistoryTableFactory struct{}

func (f *EquipmentUsageHistoryTableFactory) CreateListComponent(data any) templ.Component {
    usageHistory := data.([]db.EquipmentUsageHistory)
    return TableListEquipmentUsageHistory(usageHistory)
}

func (f *EquipmentUsageHistoryTableFactory) CreateSortComponent() templ.Component {
    return SortEquipmentUsageHistory()
}

func (f *EquipmentUsageHistoryTableFactory) GetTableName() string {
    return "equipment_usage_history"
}

// ====================
type TransportUsageHistoryTableFactory struct{}

func (f *TransportUsageHistoryTableFactory) CreateListComponent(data any) templ.Component {
    usageHistory := data.([]db.TransportUsageHistory)
    return TableListTransportUsageHistory(usageHistory)
}

func (f *TransportUsageHistoryTableFactory) CreateSortComponent() templ.Component {
    return SortTransportUsageHistory()
}

func (f *TransportUsageHistoryTableFactory) GetTableName() string {
    return "transport_usage_history"
}

// ====================
type EquipmentServiceHistoryTableFactory struct{}

func (f *EquipmentServiceHistoryTableFactory) CreateListComponent(data any) templ.Component {
    serviceHistory := data.([]db.EquipmentServiceHistory)
    return TableListEquipmentServiceHistory(serviceHistory)
}

func (f *EquipmentServiceHistoryTableFactory) CreateSortComponent() templ.Component {
    return SortEquipmentServiceHistory()
}

func (f *EquipmentServiceHistoryTableFactory) GetTableName() string {
    return "equipment_service_history"
}

// ====================
type TransportServiceHistoryTableFactory struct{}

func (f *TransportServiceHistoryTableFactory) CreateListComponent(data any) templ.Component {
    serviceHistory := data.([]db.TransportServiceHistory)
    return TableListTransportServiceHistory(serviceHistory)
}

func (f *TransportServiceHistoryTableFactory) CreateSortComponent() templ.Component {
    return SortTransportServiceHistory()
}

func (f *TransportServiceHistoryTableFactory) GetTableName() string {
    return "transport_service_history"
}

var TableFactories = map[string]TableComponentFactory{
    "equipment_types":                  &EquipmentTypesTableFactory{},
    "objects":                          &ObjectsTableFactory{},
    "accident_types":                   &AccidentTypesTableFactory{},
    "accidents":                        &AccidentsTableFactory{},
    "applications_for_admission":       &ApplicationsForAdmissionTableFactory{},
    "candidates_documents":             &CandidatesDocumentsTableFactory{},
    "candidates_medical_parameters":    &CandidatesMedicalParametersTableFactory{},
    "vgk":                              &VgkTableFactory{},
    "positions":                        &PositionsTableFactory{},
    "vgk_rescuers":                     &VgkRescuersTableFactory{},
    "vgk_rescuers_documents":           &VgkRescuersDocumentsTableFactory{},
    "vgk_locations":                    &VgkLocationsTableFactory{},
    "vgk_shifts":                       &VgkShiftsTableFactory{},
    "accidents_response_operations":    &AccidentsResponseOperationsTableFactory{},
    "operations_participations":        &OperationsParticipationsTableFactory{},
    "operations_reports":               &OperationsReportsTableFactory{},
    "trainings":                        &TrainingsTableFactory{},
    "trainings_participations":         &TrainingsParticipationsTableFactory{},
    "certifications_passings":          &CertificationsPassingsTableFactory{},
    "vgk_rescuers_medical_parameters":  &VgkRescuersMedicalParametersTableFactory{},
    "vgk_service_room":                 &VgkServiceRoomTableFactory{},
    "equipment":                        &EquipmentTableFactory{},
    "transport":                        &TransportTableFactory{},
    "equipment_usage_history":          &EquipmentUsageHistoryTableFactory{},
    "transport_usage_history":          &TransportUsageHistoryTableFactory{},
    "equipment_service_history":        &EquipmentServiceHistoryTableFactory{},
    "transport_service_history":        &TransportServiceHistoryTableFactory{},
}
