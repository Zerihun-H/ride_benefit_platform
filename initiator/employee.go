package initiator

import (
	routing "rideBenefit/internal/api/rest"
	"rideBenefit/internal/handler/rest"
	"rideBenefit/internal/module/employee"
	"rideBenefit/internal/repository"
	"rideBenefit/internal/storage/persistence"
	"rideBenefit/platform/cockroach"
	"rideBenefit/platform/httprouter"
)

// Employee initializes the domain employee
func Employee(cockroachPlatform cockroach.CockroachPlatform) []httprouter.Router {

	// Initiate the employee persistence
	employeePersistence := persistence.EmployeeInit(cockroachPlatform)
	// Initiate the employee repository
	employeeRepository := repository.EmployeeInit()
	// Initiate the employee service
	usecase := employee.Initialize(employeeRepository, employeePersistence)

	// Initiate the employee rest API handler
	handler := rest.EmployeeInit(usecase)

	// Initiate the employee routing
	return routing.EmployeeRouting(handler)

}
