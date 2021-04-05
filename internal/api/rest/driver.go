package routing

import (
	"net/http"

	"rideBenefit/internal/handler/rest"
	"rideBenefit/platform/httprouter"
)

// DriverRouting returns the list of routers for domain driver
func DriverRouting(handler rest.DriverHandler) []httprouter.Router {
	return []httprouter.Router{
		{ // Get driver
			Method:  http.MethodGet,
			Path:    "/drivers/:driverID",
			Handler: handler.GetDriver,
		}, { // Add driver
			Method:  http.MethodPost,
			Path:    "/drivers",
			Handler: handler.AddDriver,
		}, { // Update driver
			Method:  http.MethodPatch,
			Path:    "/drivers",
			Handler: handler.UpdateDriver,
		},
		{ // Delete driver
			Method:  http.MethodDelete,
			Path:    "/drivers/:driverID",
			Handler: handler.DeleteDriver,
		}, { // Bulk add drivers from excel
			Method:  http.MethodPost,
			Path:    "/drivers/excel",
			Handler: handler.AddDriversExcel,
		}, { // Bulk add drivers from CSV
			Method:  http.MethodPost,
			Path:    "/drivers/csv",
			Handler: handler.AddDriversCSV,
		},
	}
}
