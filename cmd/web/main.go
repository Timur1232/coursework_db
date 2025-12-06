package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Timur1232/coursework_db/internal/db"
	"github.com/Timur1232/coursework_db/internal/handlers"
	"github.com/Timur1232/coursework_db/views"

	"github.com/a-h/templ"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var DB *pgx.Conn

const (
	user     = "app"
	password = "app"
	host     = "127.0.0.1"
	port     = 5432
	database = "coursework"
)

func getConnString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, host, port, database)
}

func RenderPage(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func main() {
	DB, err := pgx.Connect(context.Background(), getConnString())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer DB.Close(context.Background())

	e := echo.New()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &db.DBContext{Context: c, DB: DB}
			return next(cc)
		}
	})

	e.Use(middleware.Logger())

	e.Static("/static", "static")

	e.GET("/admin", func(c echo.Context) error {
		return c.Redirect(http.StatusPermanentRedirect, "/admin/vgk")
	})

	e.GET("/admin/vgk", func(c echo.Context) error { return handlers.AdminPanel(c, "vgk", views.TableListVgk, views.SortVgk) })
	e.GET("/admin/vgk/:page", func(c echo.Context) error { return handlers.AdminPanelPage(c, "vgk", views.TableListVgk) })

	e.GET("/admin/equipment_types", func(c echo.Context) error {
		return handlers.AdminPanel(c, "equipment_types", views.TableListEquipmentTypes, views.SortEquipmentTypes)
	})
	e.GET("/admin/equipment_types/:page", func(c echo.Context) error {
		return handlers.AdminPanelPage(c, "equipment_types", views.TableListEquipmentTypes)
	})

	e.GET("/admin/objects", func(c echo.Context) error {
		return handlers.AdminPanel(c, "objects", views.TableListObjects, views.SortObjects)
	})
	e.GET("/admin/objects/:page", func(c echo.Context) error { return handlers.AdminPanelPage(c, "objects", views.TableListObjects) })

	e.GET("/admin/accident_types", func(c echo.Context) error {
		return handlers.AdminPanel(c, "accident_types", views.TableListAccidentTypes, views.SortAccidentTypes)
	})
	e.GET("/admin/accident_types/:page", func(c echo.Context) error {
		return handlers.AdminPanelPage(c, "accident_types", views.TableListAccidentTypes)
	})

	e.GET("/admin/accidents", func(c echo.Context) error {
		return handlers.AdminPanel(c, "accidents", views.TableListAccidents, views.SortAccidents)
	})
	e.GET("/admin/accidents/:page", func(c echo.Context) error { return handlers.AdminPanelPage(c, "accidents", views.TableListAccidents) })

	e.GET("/admin/applications_for_admission", func(c echo.Context) error {
		return handlers.AdminPanel(c, "applications_for_admission", views.TableListApplicationsForAdmission, views.SortApplicationsForAdmission)
	})
	e.GET("/admin/applications_for_admission/:page", func(c echo.Context) error {
		return handlers.AdminPanelPage(c, "applications_for_admission", views.TableListApplicationsForAdmission)
	})

	e.GET("/admin/candidates_documents", func(c echo.Context) error {
		return handlers.AdminPanel(c, "candidates_documents", views.TableListCandidatesDocuments, views.SortCandidatesDocuments)
	})
	e.GET("/admin/candidates_documents/:page", func(c echo.Context) error {
		return handlers.AdminPanelPage(c, "candidates_documents", views.TableListCandidatesDocuments)
	})

	e.GET("/admin/candidates_medical_parameters", func(c echo.Context) error {
		return handlers.AdminPanel(c, "candidates_medical_parameters", views.TableListCandidatesMedicalParameters, views.SortCandidatesMedicalParameters)
	})
	e.GET("/admin/candidates_medical_parameters/:page", func(c echo.Context) error {
		return handlers.AdminPanelPage(c, "candidates_medical_parameters", views.TableListCandidatesMedicalParameters)
	})

	e.GET("/admin/vgk", func(c echo.Context) error { return handlers.AdminPanel(c, "vgk", views.TableListVgk, views.SortVgk) })
	e.GET("/admin/vgk/:page", func(c echo.Context) error { return handlers.AdminPanelPage(c, "vgk", views.TableListVgk) })

	e.GET("/admin/positions", func(c echo.Context) error {
		return handlers.AdminPanel(c, "positions", views.TableListPositions, views.SortPositions)
	})
	e.GET("/admin/positions/:page", func(c echo.Context) error { return handlers.AdminPanelPage(c, "positions", views.TableListPositions) })

	e.GET("/admin/vgk_rescuers", func(c echo.Context) error {
		return handlers.AdminPanel(c, "vgk_rescuers", views.TableListVgkRescuers, views.SortVgkRescuers)
	})
	e.GET("/admin/vgk_rescuers/:page", func(c echo.Context) error {
		return handlers.AdminPanelPage(c, "vgk_rescuers", views.TableListVgkRescuers)
	})

	e.GET("/admin/vgk_rescuers_documents", func(c echo.Context) error {
		return handlers.AdminPanel(c, "vgk_rescuers_documents", views.TableListVgkRescuersDocuments, views.SortVgkRescuersDocuments)
	})
	e.GET("/admin/vgk_rescuers_documents/:page", func(c echo.Context) error {
		return handlers.AdminPanelPage(c, "vgk_rescuers_documents", views.TableListVgkRescuersDocuments)
	})

	e.GET("/admin/vgk_locations", func(c echo.Context) error {
		return handlers.AdminPanel(c, "vgk_locations", views.TableListVgkLocations, views.SortVgkLocations)
	})
	e.GET("/admin/vgk_locations/:page", func(c echo.Context) error {
		return handlers.AdminPanelPage(c, "vgk_locations", views.TableListVgkLocations)
	})

	e.GET("/admin/vgk_shifts", func(c echo.Context) error {
		return handlers.AdminPanel(c, "vgk_shifts", views.TableListVgkShifts, views.SortVgkShifts)
	})
	e.GET("/admin/vgk_shifts/:page", func(c echo.Context) error { return handlers.AdminPanelPage(c, "vgk_shifts", views.TableListVgkShifts) })

	e.GET("/admin/accidents_response_operations", func(c echo.Context) error {
		return handlers.AdminPanel(c, "accidents_response_operations", views.TableListAccidentsResponseOperations, views.SortAccidentsResponseOperations)
	})
	e.GET("/admin/accidents_response_operations/:page", func(c echo.Context) error {
		return handlers.AdminPanelPage(c, "accidents_response_operations", views.TableListAccidentsResponseOperations)
	})

	e.GET("/admin/operations_participations", func(c echo.Context) error {
		return handlers.AdminPanel(c, "operations_participations", views.TableListOperationsParticipations, views.SortOperationsParticipations)
	})
	e.GET("/admin/operations_participations/:page", func(c echo.Context) error {
		return handlers.AdminPanelPage(c, "operations_participations", views.TableListOperationsParticipations)
	})

	e.GET("/admin/operations_reports", func(c echo.Context) error {
		return handlers.AdminPanel(c, "operations_reports", views.TableListOperationsReports, views.SortOperationsReports)
	})
	e.GET("/admin/operations_reports/:page", func(c echo.Context) error {
		return handlers.AdminPanelPage(c, "operations_reports", views.TableListOperationsReports)
	})

	e.GET("/admin/trainings", func(c echo.Context) error {
		return handlers.AdminPanel(c, "trainings", views.TableListTrainings, views.SortTrainings)
	})
	e.GET("/admin/trainings/:page", func(c echo.Context) error { return handlers.AdminPanelPage(c, "trainings", views.TableListTrainings) })

	e.GET("/admin/trainings_participations", func(c echo.Context) error {
		return handlers.AdminPanel(c, "trainings_participations", views.TableListTrainingsParticipations, views.SortTrainingsParticipations)
	})
	e.GET("/admin/trainings_participations/:page", func(c echo.Context) error {
		return handlers.AdminPanelPage(c, "trainings_participations", views.TableListTrainingsParticipations)
	})

	e.GET("/admin/certifications_passings", func(c echo.Context) error {
		return handlers.AdminPanel(c, "certifications_passings", views.TableListCertificationsPassings, views.SortCertificationsPassings)
	})
	e.GET("/admin/certifications_passings/:page", func(c echo.Context) error {
		return handlers.AdminPanelPage(c, "certifications_passings", views.TableListCertificationsPassings)
	})

	e.GET("/admin/vgk_rescuers_medical_parameters", func(c echo.Context) error {
		return handlers.AdminPanel(c, "vgk_rescuers_medical_parameters", views.TableListVgkRescuersMedicalParameters, views.SortVgkRescuersMedicalParameters)
	})
	e.GET("/admin/vgk_rescuers_medical_parameters/:page", func(c echo.Context) error {
		return handlers.AdminPanelPage(c, "vgk_rescuers_medical_parameters", views.TableListVgkRescuersMedicalParameters)
	})

	e.GET("/admin/vgk_service_room", func(c echo.Context) error {
		return handlers.AdminPanel(c, "vgk_service_room", views.TableListVgkServiceRoom, views.SortVgkServiceRoom)
	})
	e.GET("/admin/vgk_service_room/:page", func(c echo.Context) error {
		return handlers.AdminPanelPage(c, "vgk_service_room", views.TableListVgkServiceRoom)
	})

	e.GET("/admin/equipment", func(c echo.Context) error {
		return handlers.AdminPanel(c, "equipment", views.TableListEquipment, views.SortEquipment)
	})
	e.GET("/admin/equipment/:page", func(c echo.Context) error { return handlers.AdminPanelPage(c, "equipment", views.TableListEquipment) })

	e.GET("/admin/transport", func(c echo.Context) error {
		return handlers.AdminPanel(c, "transport", views.TableListTransport, views.SortTransport)
	})
	e.GET("/admin/transport/:page", func(c echo.Context) error { return handlers.AdminPanelPage(c, "transport", views.TableListTransport) })

	e.GET("/admin/equipment_usage_history", func(c echo.Context) error {
		return handlers.AdminPanel(c, "equipment_usage_history", views.TableListEquipmentUsageHistory, views.SortEquipmentUsageHistory)
	})
	e.GET("/admin/equipment_usage_history/:page", func(c echo.Context) error {
		return handlers.AdminPanelPage(c, "equipment_usage_history", views.TableListEquipmentUsageHistory)
	})

	e.GET("/admin/transport_usage_history", func(c echo.Context) error {
		return handlers.AdminPanel(c, "transport_usage_history", views.TableListTransportUsageHistory, views.SortTransportUsageHistory)
	})
	e.GET("/admin/transport_usage_history/:page", func(c echo.Context) error {
		return handlers.AdminPanelPage(c, "transport_usage_history", views.TableListTransportUsageHistory)
	})

	e.GET("/admin/equipment_service_history", func(c echo.Context) error {
		return handlers.AdminPanel(c, "equipment_service_history", views.TableListEquipmentServiceHistory, views.SortEquipmentServiceHistory)
	})
	e.GET("/admin/equipment_service_history/:page", func(c echo.Context) error {
		return handlers.AdminPanelPage(c, "equipment_service_history", views.TableListEquipmentServiceHistory)
	})

	e.GET("/admin/transport_service_history", func(c echo.Context) error {
		return handlers.AdminPanel(c, "transport_service_history", views.TableListTransportServiceHistory, views.SortTransportServiceHistory)
	})
	e.GET("/admin/transport_service_history/:page", func(c echo.Context) error {
		return handlers.AdminPanelPage(c, "transport_service_history", views.TableListTransportServiceHistory)
	})

	e.Logger.Fatal(e.Start(":42069"))

}
