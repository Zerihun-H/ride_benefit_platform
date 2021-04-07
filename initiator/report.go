package initiator

import (
	routing "rideBenefit/internal/api/rest"
	"rideBenefit/internal/handler/rest"
	"rideBenefit/internal/module/report"
	"rideBenefit/internal/repository"
	"rideBenefit/internal/storage/persistence"
	"rideBenefit/platform/cockroach"
	"rideBenefit/platform/httprouter"
)

// Report initializes the domain report
func Report(cockroachPlatform cockroach.CockroachPlatform) []httprouter.Router {

	// Initiate the report persistence
	reportPersistence := persistence.ReportInit(cockroachPlatform)
	// Initiate the report repository
	reportRepository := repository.ReportInit()
	// Initiate the report service
	usecase := report.Initialize(reportRepository, reportPersistence)

	// Initiate the report rest API handler
	handler := rest.ReportInit(usecase)

	// Initiate the report routing
	return routing.ReportRouting(handler)

}
