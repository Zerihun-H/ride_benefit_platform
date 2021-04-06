package repository

// EmployeeRepository contains the functions of data logic for domain employee
type EmployeeRepository interface {
}

type employeeRepository struct {
}

// EmployeeInit initializes the data logic / repository for domain employee
func EmployeeInit() EmployeeRepository {
	return employeeRepository{}
}
