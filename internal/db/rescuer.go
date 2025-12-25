package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

func GetRescuerOperationsHistory(DB *pgx.Conn, rescuerID uint64) ([]OperationsParticipations, error) {
	query := `
		SELECT op.*, aro.start_date_time, aro.end_date_time, aro.status as operation_status,
		       acc.accident_type, acc.begin_date_time as accident_start
		FROM operations_participations op
		JOIN accidents_response_operations aro ON op.id_operation = aro.id_operation
		JOIN accidents acc ON aro.id_accident = acc.id_accident
		JOIN vgk v ON op.id_vgk = v.id_vgk
		JOIN vgk_rescuers vr ON v.id_vgk = vr.id_vgk
		WHERE vr.id_rescuer = $1
		ORDER BY aro.start_date_time DESC
	`

	type OperationHistory struct {
		IdVgk           uint64     `db:"id_vgk"`
		IdOperation     uint64     `db:"id_operation"`
		AssignedTask    string     `db:"assigned_task"`
		StartDateTime   time.Time  `db:"start_date_time"`
		EndDateTime     *time.Time `db:"end_date_time"`
		OperationStatus string     `db:"operation_status"`
		AccidentType    string     `db:"accident_type"`
		AccidentStart   time.Time  `db:"accident_start"`
	}

	result, err := Query[OperationHistory](DB, context.Background(), query, rescuerID)
	if err != nil {
		return nil, err
	}

	history := result.([]OperationHistory)
	var participations []OperationsParticipations
	for _, h := range history {
		participations = append(participations, OperationsParticipations{
			IdVgk:       h.IdVgk,
			IdOperation: h.IdOperation,
			AssignedTask: fmt.Sprintf("%s (Авария: %s, Начало: %s)",
				h.AssignedTask, h.AccidentType, h.AccidentStart.Format("02.01.2006")),
		})
	}

	return participations, nil
}

func GetRescuerShiftsHistory(DB *pgx.Conn, rescuerID uint64) ([]VgkShifts, error) {
	query := `
		SELECT vs.*, vl.address as location_address
		FROM vgk_shifts vs
		JOIN vgk_locations vl ON vs.id_vgk_location = vl.id_vgk_location
		JOIN vgk_rescuers vr ON vs.id_vgk = vr.id_vgk
		WHERE vr.id_rescuer = $1
		ORDER BY vs.shift_start DESC
	`

	type ShiftWithLocation struct {
		VgkShifts
		LocationAddress string `db:"location_address"`
	}

	result, err := Query[ShiftWithLocation](DB, context.Background(), query, rescuerID)
	if err != nil {
		return nil, err
	}

	shiftsWithLoc := result.([]ShiftWithLocation)
	var shifts []VgkShifts
	for _, s := range shiftsWithLoc {
		shift := s.VgkShifts
		shifts = append(shifts, shift)
	}

	return shifts, nil
}

func GetRescuerCurrentOperations(DB *pgx.Conn, rescuerID uint64) ([]OperationsParticipations, error) {
	query := `
		SELECT op.*, aro.start_date_time, aro.end_date_time, aro.status as operation_status,
		       acc.accident_type, acc.begin_date_time as accident_start
		FROM operations_participations op
		JOIN accidents_response_operations aro ON op.id_operation = aro.id_operation
		JOIN accidents acc ON aro.id_accident = acc.id_accident
		JOIN vgk v ON op.id_vgk = v.id_vgk
		JOIN vgk_rescuers vr ON v.id_vgk = vr.id_vgk
		WHERE vr.id_rescuer = $1 
		AND aro.status NOT IN ('completed', 'failed')
		AND (aro.end_date_time IS NULL OR aro.end_date_time > NOW())
		ORDER BY aro.start_date_time DESC
	`

	type CurrentOperation struct {
		IdVgk           uint64     `db:"id_vgk"`
		IdOperation     uint64     `db:"id_operation"`
		AssignedTask    string     `db:"assigned_task"`
		StartDateTime   time.Time  `db:"start_date_time"`
		EndDateTime     *time.Time `db:"end_date_time"`
		OperationStatus string     `db:"operation_status"`
		AccidentType    string     `db:"accident_type"`
		AccidentStart   time.Time  `db:"accident_start"`
	}

	result, err := Query[CurrentOperation](DB, context.Background(), query, rescuerID)
	if err != nil {
		return nil, err
	}

	currentOps := result.([]CurrentOperation)
	var participations []OperationsParticipations
	for _, op := range currentOps {
		participations = append(participations, OperationsParticipations{
			IdVgk:       op.IdVgk,
			IdOperation: op.IdOperation,
			AssignedTask: fmt.Sprintf("%s (Авария: %s, Статус: %s)",
				op.AssignedTask, op.AccidentType, op.OperationStatus),
		})
	}

	return participations, nil
}

func GetRescuerCurrentShifts(DB *pgx.Conn, rescuerID uint64) ([]VgkShifts, error) {
	query := `
		SELECT vs.*, vl.address as location_address
		FROM vgk_shifts vs
		JOIN vgk_locations vl ON vs.id_vgk_location = vl.id_vgk_location
		JOIN vgk_rescuers vr ON vs.id_vgk = vr.id_vgk
		WHERE vr.id_rescuer = $1 
		AND vs.shift_start <= NOW() 
		AND vs.shift_end >= NOW()
		ORDER BY vs.shift_start DESC
	`

	type CurrentShift struct {
		VgkShifts
		LocationAddress string `db:"location_address"`
	}

	result, err := Query[CurrentShift](DB, context.Background(), query, rescuerID)
	if err != nil {
		return nil, err
	}

	currentShifts := result.([]CurrentShift)
	var shifts []VgkShifts
	for _, s := range currentShifts {
		shift := s.VgkShifts
		shifts = append(shifts, shift)
	}

	return shifts, nil
}

func GetTeamVGKLocations(DB *pgx.Conn, vgkID uint64) ([]VgkLocations, error) {
	query := `
		SELECT * FROM vgk_locations 
		WHERE id_vgk_location IN (
			SELECT DISTINCT id_vgk_location FROM vgk_shifts WHERE id_vgk = $1
		) OR id_responsible IN (
			SELECT id_rescuer FROM vgk_rescuers WHERE id_vgk = $1
		)
		ORDER BY address
	`

	result, err := Query[VgkLocations](DB, context.Background(), query, vgkID)
	if err != nil {
		return nil, err
	}

	return result.([]VgkLocations), nil
}

func GetTeamServiceRooms(DB *pgx.Conn, vgkID uint64) ([]VgkServiceRoom, error) {
	query := `
		SELECT * FROM vgk_service_room 
		WHERE id_responsible IN (
			SELECT id_rescuer FROM vgk_rescuers WHERE id_vgk = $1
		)
		ORDER BY purpose
	`

	result, err := Query[VgkServiceRoom](DB, context.Background(), query, vgkID)
	if err != nil {
		return nil, err
	}

	return result.([]VgkServiceRoom), nil
}
