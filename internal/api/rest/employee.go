package routing

import (
	"net/http"

	"rideBenefit/internal/handler/rest"
	"rideBenefit/platform/httprouter"
)

// EmployeeRouting returns the list of routers for domain employee
func EmployeeRouting(handler rest.EmployeeHandler) []httprouter.Router {
	return []httprouter.Router{
		{ // Get employee
			Method:  http.MethodGet,
			Path:    "/employees/:employeeID",
			Handler: handler.GetEmployee,
		}, { // Add employee
			Method:  http.MethodPost,
			Path:    "/employees",
			Handler: handler.AddEmployee,
		}, { // Update employee
			Method:  http.MethodPatch,
			Path:    "/employees",
			Handler: handler.UpdateEmployee,
		},
		{ // Delete employee
			Method:  http.MethodDelete,
			Path:    "/employees/:employeeID",
			Handler: handler.DeleteEmployee,
		}, { // Bulk add employees from excel
			Method:  http.MethodPost,
			Path:    "/employees/excel",
			Handler: handler.AddEmployeesExcel,
		}, { // Bulk add employees from CSV
			Method:  http.MethodPost,
			Path:    "/employees/csv",
			Handler: handler.AddEmployeesCSV,
		},
	}
}
