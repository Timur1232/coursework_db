package handlers

import (
	"context"
	"fmt"

	"github.com/Timur1232/coursework_db/internal/db"
	"github.com/Timur1232/coursework_db/views"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Profile(c echo.Context) error {
	user := c.(*db.DBContext).User

	var profileContent templ.Component

	if user == nil {
		profileContent = views.ProfileGuest(user)
	} else {
		switch user.Role {
		case db.Role_Candidate:
			application, err := db.GetApplicationByUserID(c.(*db.DBContext).DB, user.IdUser)
			if err != nil {
				profileContent = views.ProfileGuest(user)
				break
			}

			documents, _ := db.GetCandidateDocuments(c.(*db.DBContext).DB, application.IdApplication)
			medicalParams, _ := db.GetCandidateMedicalParameters(c.(*db.DBContext).DB, application.IdApplication)
			profileContent = views.ProfileCandidate(user, application, documents, medicalParams)

		case db.Role_Rescuer:
			rescuer, err := db.GetRescuerByUserID(c.(*db.DBContext).DB, user.IdUser)
			if err != nil {

				profileContent = views.ProfileBasic(user)
				break
			}

			documents, _ := db.GetRescuerDocuments(c.(*db.DBContext).DB, rescuer.IdRescuer)

			medicalParams, _ := db.GetRescuerMedicalParameters(c.(*db.DBContext).DB, rescuer.IdRescuer)

			var teamMembers []db.VgkRescuers
			var vgkDetails *db.Vgk
			var operationsHistory []db.OperationsParticipations
			var shiftsHistory []db.VgkShifts
			var currentOperations []db.OperationsParticipations
			var currentShifts []db.VgkShifts
			var vgkLocations []db.VgkLocations
			var serviceRooms []db.VgkServiceRoom

			if rescuer.IdVgk.Valid {
				teamMembers, _ = db.GetTeamMembers(c.(*db.DBContext).DB, uint64(rescuer.IdVgk.Int64))
				vgkDetails, _ = db.GetVGKDetails(c.(*db.DBContext).DB, uint64(rescuer.IdVgk.Int64))

				operationsHistory, err = db.GetRescuerOperationsHistory(c.(*db.DBContext).DB, rescuer.IdRescuer)
				if err != nil {
					fmt.Println("operations: ", err)
				}
				fmt.Println(operationsHistory)
				shiftsHistory, err = db.GetRescuerShiftsHistory(c.(*db.DBContext).DB, rescuer.IdRescuer)
				if err != nil {
					fmt.Println("shifts: ", err)
				}
				fmt.Println(shiftsHistory)
				currentOperations, err = db.GetRescuerCurrentOperations(c.(*db.DBContext).DB, rescuer.IdRescuer)
				if err != nil {
					fmt.Println("current operation: ", err)
				}
				fmt.Println(currentOperations)
				currentShifts, err = db.GetRescuerCurrentShifts(c.(*db.DBContext).DB, rescuer.IdRescuer)
				if err != nil {
					fmt.Println("current shift: ", err)
				}
				fmt.Println(currentShifts)
				vgkLocations, err = db.GetTeamVGKLocations(c.(*db.DBContext).DB, uint64(rescuer.IdVgk.Int64))
				if err != nil {
					fmt.Println("vgk locations: ", err)
				}
				fmt.Println(vgkLocations)
				serviceRooms, err = db.GetTeamServiceRooms(c.(*db.DBContext).DB, uint64(rescuer.IdVgk.Int64))
				if err != nil {
					fmt.Println("sercice rooms: ", err)
				}
				fmt.Println(serviceRooms)

			}

			profileContent = views.ProfileRescuer(user, rescuer, documents, medicalParams,
				vgkDetails, teamMembers, operationsHistory, shiftsHistory,
				currentOperations, currentShifts, vgkLocations, serviceRooms)
		case db.Role_Operator:
			profileContent = views.ProfileOperator(user)

		case db.Role_Admin:
			profileContent = views.ProfileAdmin(user)

		case db.Role_Guest:
			profileContent = views.ProfileGuest(user)

		default:
			profileContent = views.ProfileBasic(user)
		}
	}

	if c.Request().Header.Get("HX-Request") == "true" {
		return RenderPage(c, profileContent)
	}
	page := views.Layout("Личный кабинет", profileContent, user)
	return RenderPage(c, page)
}

func CancelApplication(c echo.Context) error {
	user := c.(*db.DBContext).User

	if user.Role != db.Role_Candidate {
		return c.Redirect(303, "/profile")
	}

	application, err := db.GetApplicationByUserID(c.(*db.DBContext).DB, user.IdUser)
	if err != nil {
		return c.Redirect(303, "/profile")
	}

	err = db.DeleteApplication(c.(*db.DBContext).DB, application.IdApplication)
	if err != nil {
		return c.String(500, "Ошибка при удалении заявления")
	}

	_, err = c.(*db.DBContext).DB.Exec(context.Background(),
		"UPDATE users SET role = $1 WHERE id_user = $2",
		db.Role_Guest, user.IdUser)
	if err != nil {
		return c.String(500, "Ошибка при обновлении роли")
	}

	user.Role = db.Role_Guest

	profileContent := views.ProfileGuest(user)

	if c.Request().Header.Get("HX-Request") == "true" {
		return RenderPage(c, profileContent)
	}
	page := views.Layout("Личный кабинет", profileContent, user)
	return RenderPage(c, page)
}
