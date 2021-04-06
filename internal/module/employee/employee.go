package employee

import (
	"rideBenefit/internal/constant/model"
)

func (s *service) GetEmployee(employeeID uint64) (*model.Employee, error) {
	// Some validation

	return s.employeePersist.GetEmployee(employeeID)
}

func (s *service) AddEmployee(employee *model.Employee) (*model.Employee, error) {

	return s.employeePersist.AddEmployee(employee)
}

func (s *service) UpdateEmployee(employee *model.Employee) (*model.Employee, error) {

	return s.employeePersist.UpdateEmployee(employee)

}

func (s *service) DeleteEmployee(employeeID uint64) error {

	return s.employeePersist.DeleteEmployee(employeeID)

}

func (s *service) AddEmployees(employees []model.Employee) error {

	return s.employeePersist.AddEmployees(employees)
}
