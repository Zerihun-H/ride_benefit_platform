package persistence

import (
	"rideBenefit/internal/constant/model"
	"rideBenefit/platform/cockroach"
)

// EmployeePersistence contains the list of functions for database table employees
type EmployeePersistence interface {
	GetEmployee(employeeID uint64) (*model.Employee, error)
	AddEmployee(employee *model.Employee) (*model.Employee, error)
	UpdateEmployee(employee *model.Employee) (*model.Employee, error)
	DeleteEmployee(employeeID uint64) error
	AddEmployees(employees []model.Employee) error
}

type employeePersistence struct {
	db cockroach.CockroachPlatform
}

// EmployeeInit is to init the employee persistence that contains employee data
func EmployeeInit(db cockroach.CockroachPlatform) EmployeePersistence {
	return &employeePersistence{
		db,
	}
}

// Getemployee using the employee id fetchs the employee from the employee database
func (dp *employeePersistence) GetEmployee(employeeID uint64) (*model.Employee, error) {
	db, err := dp.db.Open()
	if err != nil {
		return nil, err
	}
	dbc, err := db.DB()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()

	employee := &model.Employee{}
	if err := db.Where("id = ?", employeeID).First(employee).Error; err != nil {

		return &model.Employee{}, err
	}
	return employee, nil
}

// AddEmployee is adds a employee to the database given a valid diver
func (dp *employeePersistence) AddEmployee(employee *model.Employee) (*model.Employee, error) {
	db, err := dp.db.Open()
	if err != nil {
		return nil, err
	}
	dbc, err := db.DB()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()
	if err := db.Create(employee).Error; err != nil {
		return nil, err
	}
	return employee, nil
}

// AddEmployee is adds a employee to the database given a valid diver
func (dp *employeePersistence) UpdateEmployee(employee *model.Employee) (*model.Employee, error) {
	db, err := dp.db.Open()
	if err != nil {
		return nil, err
	}
	dbc, err := db.DB()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()
	updatedEmployee := employee
	err = db.First(employee).Error
	if err != nil {
		return nil, err
	}
	updatedEmployee.ID = employee.ID
	err = db.Save(&updatedEmployee).Error
	if err != nil {
		return nil, err
	}
	return employee, nil
}

// DeleteEmployee is adds a employee to the database given a valid diver
func (dp *employeePersistence) DeleteEmployee(employeeID uint64) error {
	db, err := dp.db.Open()
	if err != nil {
		return err
	}
	dbc, err := db.DB()
	if err != nil {
		return err
	}
	defer dbc.Close()
	return db.Where("id = ?", employeeID).Delete(&model.Employee{}).Error
}

// AddEmployees is adds a employees to the database given valid divers
func (dp *employeePersistence) AddEmployees(employees []model.Employee) error {
	db, err := dp.db.Open()
	if err != nil {
		return err
	}
	dbc, err := db.DB()
	if err != nil {
		return err
	}
	defer dbc.Close()
	return db.Create(employees).Error
}
