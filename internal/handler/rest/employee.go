package rest

import (
	"encoding/csv"
	"encoding/json"
	"net/http"
	"rideBenefit/internal/constant"
	model "rideBenefit/internal/constant/model"
	"rideBenefit/internal/module/auth"
	"rideBenefit/internal/module/employee"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

const (
	birthdateLayout = "2006-01-02T15:04:05.000000Z"
	sheetName       = "Sheet1"
)

// EmployeeHandler contains the function of handler for domain Employee
type EmployeeHandler interface {
	GetEmployee(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	AddEmployee(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	UpdateEmployee(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	DeleteEmployee(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	AddEmployeesExcel(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	AddEmployeesCSV(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type employeeHandler struct {
	Usecase employee.Usecase
}

// EmployeeInit is to initialize the rest handler for domain Employee
func EmployeeInit(Usecase employee.Usecase) EmployeeHandler {

	return &employeeHandler{
		Usecase,
	}
}

func (dh *employeeHandler) GetEmployee(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	roleID, err := strconv.Atoi(r.Header.Get("rle"))
	if err != nil {
		http.Error(w, constant.ErrInvalidEmployeeID.Error(), http.StatusBadRequest)
		return
	}
	// Check if the role has the required permission
	yes, err := auth.AuthService().RoleHasPermission(uint64(roleID), constant.PermFetchEmployee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !yes {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Check if the employee ID param is valid

	employeeID := ps.ByName("employeeID")
	// Convert the employeeID string to uint64
	id, err := strconv.Atoi(employeeID)
	if err != nil {
		http.Error(w, constant.ErrInvalidEmployeeID.Error(), http.StatusBadRequest)
		return
	}

	employee, err := dh.Usecase.GetEmployee(uint64(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, gorm.ErrRecordNotFound.Error(), http.StatusNotFound)
			return
		} else {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employee)

}

func (dh *employeeHandler) AddEmployee(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse employee data
	employee := &model.Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		http.Error(w, constant.ErrInvalidRequestBody.Error(), http.StatusBadRequest)
		return
	}
	drv, err := dh.Usecase.AddEmployee(employee)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(drv)

}

func (dh *employeeHandler) UpdateEmployee(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse employee data
	employee := &model.Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		http.Error(w, constant.ErrInvalidRequestBody.Error(), http.StatusBadRequest)
		return
	}

	drv, err := dh.Usecase.UpdateEmployee(employee)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(drv)
}

func (dh *employeeHandler) DeleteEmployee(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Check if the diver ID param is valid
	employeeID := ps.ByName("employeeID")
	// Convert the employeeID string to uint64
	id, err := strconv.Atoi(employeeID)
	if err != nil {
		http.Error(w, constant.ErrInvalidEmployeeID.Error(), http.StatusBadRequest)
	}
	err = dh.Usecase.DeleteEmployee(uint64(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, constant.ErrInvalidEmployeeID.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (dh *employeeHandler) AddEmployeesExcel(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	file, _, err := r.FormFile("employees")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	f, err := excelize.OpenReader(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	employees, err := parseEmployeesExcel(f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = dh.Usecase.AddEmployees(employees)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (dh *employeeHandler) AddEmployeesCSV(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	file, _, err := r.FormFile("employees")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	csv := csv.NewReader(file)

	employees, err := parseEmployeesCSV(csv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = dh.Usecase.AddEmployees(employees)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func parseEmployeesExcel(f *excelize.File) ([]model.Employee, error) {

	rows, err := f.GetRows(sheetName)
	if err != nil {
		return []model.Employee{}, err
	}

	employees := []model.Employee{}

	for i, row := range rows {

		if i == 0 {
			continue
		}

		employee := model.Employee{}

		// Parse EmployeeID
		employeeID, err := strconv.Atoi(row[0])
		if err != nil {
			return []model.Employee{}, err
		}

		employee.ID = uint64(employeeID)
		employee.FirstName = row[1]
		employee.LastName = row[2]
		employee.Surname = row[3]
		// Parse date
		birthdate, err := time.Parse(birthdateLayout, row[4])

		if err != nil {
			return []model.Employee{}, err
		}

		employee.BirthDate = &birthdate

		employee.Gender = row[5]
		employee.Email = row[6]
		employee.PhoneNumber = row[7]
		employee.EmergencyContact = row[8]
		employee.EmergencyNumber = row[9]
		// parse age
		age, err := strconv.Atoi(row[10])
		if err != nil {

			return []model.Employee{}, err
		}
		employee.Age = uint32(age)
		employee.Type = row[11]
		employees = append(employees, employee)
	}
	return employees, nil
}

func parseEmployeesCSV(cr *csv.Reader) ([]model.Employee, error) {

	records, err := cr.ReadAll()
	if err != nil {
		return []model.Employee{}, err
	}
	employees := []model.Employee{}
	for i, record := range records {
		if i == 0 {
			continue
		}

		employee := model.Employee{}
		// Parse EmployeeID
		employeeID, err := strconv.Atoi(record[0])
		if err != nil {
			return []model.Employee{}, err
		}
		employee.ID = uint64(employeeID)
		employee.FirstName = record[1]
		employee.LastName = record[2]
		employee.Surname = record[3]
		// Parse date
		birthdate, err := time.Parse(birthdateLayout, record[4])

		if err != nil {
			return []model.Employee{}, err
		}
		employee.BirthDate = &birthdate

		employee.Gender = record[5]
		employee.Email = record[6]
		employee.PhoneNumber = record[7]
		employee.EmergencyContact = record[8]
		employee.EmergencyNumber = record[9]
		// parse age
		age, err := strconv.Atoi(record[10])
		if err != nil {
			return []model.Employee{}, err
		}
		employee.Age = uint32(age)
		employee.Type = record[11]
		employees = append(employees, employee)

	}

	return employees, nil
}
