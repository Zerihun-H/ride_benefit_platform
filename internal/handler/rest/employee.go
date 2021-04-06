package rest

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rideBenefit/internal/constant/model"
	Employee "rideBenefit/internal/module/employee"
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
	EmployeeCase Employee.Usecase
}

// EmployeeInit is to initialize the rest handler for domain Employee
func EmployeeInit(EmployeeCase Employee.Usecase) EmployeeHandler {
	return &employeeHandler{
		EmployeeCase,
	}
}

func (dh *employeeHandler) GetEmployee(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Check if the diver ID param is valid
	employeeID := ps.ByName("employeeID")
	// Convert the employeeID string to uint64
	id, err := strconv.Atoi(employeeID)
	if err != nil {
		http.Error(w, model.ErrInvalidEmployeeID.Error(), http.StatusBadRequest)
		return
	}

	employee, err := dh.EmployeeCase.GetEmployee(uint64(id))
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
		http.Error(w, model.ErrInvalidRequestBody.Error(), http.StatusBadRequest)
		return
	}
	drv, err := dh.EmployeeCase.AddEmployee(employee)
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
		http.Error(w, model.ErrInvalidRequestBody.Error(), http.StatusBadRequest)
		return
	}

	drv, err := dh.EmployeeCase.UpdateEmployee(employee)
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
		http.Error(w, model.ErrInvalidEmployeeID.Error(), http.StatusBadRequest)
	}
	err = dh.EmployeeCase.DeleteEmployee(uint64(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, model.ErrInvalidEmployeeID.Error(), http.StatusBadRequest)
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

	err = dh.EmployeeCase.AddEmployees(employees)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (dh *employeeHandler) AddEmployeesCSV(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	file, header, err := r.FormFile("employees")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("File Name", header.Filename)
	csv := csv.NewReader(file)

	employees, err := parseEmployeesCSV(csv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = dh.EmployeeCase.AddEmployees(employees)
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
			log.Println("parsing id", err)
			log.Println("id ", row[0])

			return []model.Employee{}, err
		}

		employee.ID = uint64(employeeID)
		employee.FirstName = row[1]
		employee.LastName = row[2]
		employee.Surname = row[3]
		// Parse date
		birthdate, err := time.Parse(birthdateLayout, row[4])

		if err != nil {
			log.Println("parsing date", err)
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
			log.Println("parsing age", err)

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
		fmt.Printf("FirstName: %s LastName %s\n", record[1], record[2])

		employee := model.Employee{}
		// Parse EmployeeID
		employeeID, err := strconv.Atoi(record[0])
		if err != nil {
			panic(err)
		}
		employee.ID = uint64(employeeID)
		employee.FirstName = record[1]
		employee.LastName = record[2]
		employee.Surname = record[3]
		// Parse date
		birthdate, err := time.Parse(birthdateLayout, record[4])

		if err != nil {
			fmt.Println(err)
		}
		employee.BirthDate = &birthdate

		employee.Gender = record[5]
		employee.Email = record[6]
		employee.PhoneNumber = record[7]
		employee.EmergencyContact = record[8]
		employee.EmergencyNumber = record[9]
		employee.Type = record[10]
		// parse age
		age, err := strconv.Atoi(record[11])
		if err != nil {
			panic(err)
		}
		employee.Age = uint32(age)
		employees = append(employees, employee)

	}

	return employees, nil
}
