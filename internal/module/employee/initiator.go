package employee

import (
	"rideBenefit/internal/constant/model"
	"rideBenefit/internal/repository"
	"rideBenefit/internal/storage/persistence"
)

// Usecase contains the function of business logic of domain employee
type Usecase interface {
	GetEmployee(employeeID uint64) (*model.Employee, error)
	AddEmployee(employee *model.Employee) (*model.Employee, error)
	UpdateEmployee(employee *model.Employee) (*model.Employee, error)
	DeleteEmployee(employeeID uint64) error
	AddEmployees(employees []model.Employee) error
}

type service struct {
	employeeRepo    repository.EmployeeRepository
	employeePersist persistence.EmployeePersistence
}

// Initialize takes all necessary service for domain employee to run the business logic of domain employee
func Initialize(employeeRepo repository.EmployeeRepository, employeePersist persistence.EmployeePersistence) Usecase {
	return &service{
		employeeRepo:    employeeRepo,
		employeePersist: employeePersist,
	}
}
