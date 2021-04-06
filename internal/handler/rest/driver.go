package rest

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rideBenefit/internal/constant/model"
	Driver "rideBenefit/internal/module/driver"
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

// DriverHandler contains the function of handler for domain Driver
type DriverHandler interface {
	GetDriver(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	AddDriver(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	UpdateDriver(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	DeleteDriver(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	AddDriversExcel(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	AddDriversCSV(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type driverHandler struct {
	DriverCase Driver.Usecase
}

// DriverInit is to initialize the rest handler for domain Driver
func DriverInit(DriverCase Driver.Usecase) DriverHandler {
	return &driverHandler{
		DriverCase,
	}
}

func (dh *driverHandler) GetDriver(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Check if the diver ID param is valid
	driverID := ps.ByName("driverID")
	// Convert the driverID string to uint64
	id, err := strconv.Atoi(driverID)
	if err != nil {
		http.Error(w, model.ErrInvalidDriverID.Error(), http.StatusBadRequest)
		return
	}

	driver, err := dh.DriverCase.GetDriver(uint64(id))
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
	json.NewEncoder(w).Encode(driver)

}

func (dh *driverHandler) AddDriver(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse driver data
	driver := &model.Driver{}
	err := json.NewDecoder(r.Body).Decode(&driver)
	if err != nil {
		http.Error(w, model.ErrInvalidRequestBody.Error(), http.StatusBadRequest)
		return
	}
	drv, err := dh.DriverCase.AddDriver(driver)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(drv)

}
func (dh *driverHandler) UpdateDriver(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse driver data
	driver := &model.Driver{}
	err := json.NewDecoder(r.Body).Decode(&driver)
	if err != nil {
		http.Error(w, model.ErrInvalidRequestBody.Error(), http.StatusBadRequest)
		return
	}

	drv, err := dh.DriverCase.UpdateDriver(driver)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(drv)
}

func (dh *driverHandler) DeleteDriver(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Check if the diver ID param is valid
	driverID := ps.ByName("driverID")
	// Convert the driverID string to uint64
	id, err := strconv.Atoi(driverID)
	if err != nil {
		http.Error(w, model.ErrInvalidDriverID.Error(), http.StatusBadRequest)
	}
	err = dh.DriverCase.DeleteDriver(uint64(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, model.ErrInvalidDriverID.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (dh *driverHandler) AddDriversExcel(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	file, header, err := r.FormFile("drivers")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	f, err := excelize.OpenReader(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("File Name", header.Filename)

	drivers, err := parseDriversExcel(f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = dh.DriverCase.AddDrivers(drivers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (dh *driverHandler) AddDriversCSV(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	file, header, err := r.FormFile("drivers")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("File Name", header.Filename)
	csv := csv.NewReader(file)

	drivers, err := parseDrivesCSV(csv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = dh.DriverCase.AddDrivers(drivers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func parseDriversExcel(f *excelize.File) ([]model.Driver, error) {

	rows, err := f.GetRows(sheetName)
	if err != nil {
		return []model.Driver{}, err
	}

	drivers := []model.Driver{}
	for i, row := range rows {
		if i == 0 {
			continue
		}
		driver := model.Driver{}
		// Parse DriverID
		driverID, err := strconv.Atoi(row[0])
		if err != nil {
			return []model.Driver{}, err
		}

		driver.ID = uint64(driverID)
		driver.FirstName = row[1]
		driver.LastName = row[2]
		driver.Surname = row[3]
		// Parse date
		birthdate, err := time.Parse(birthdateLayout, row[4])

		if err != nil {
			return []model.Driver{}, err
		}

		driver.BirthDate = &birthdate

		driver.Gender = row[5]
		driver.Email = row[6]
		driver.PhoneNumber = row[7]
		driver.SideNumber = row[8]
		driver.EmergencyContact = row[9]
		driver.EmergencyNumber = row[10]
		// parse age
		age, err := strconv.Atoi(row[11])
		if err != nil {
			return []model.Driver{}, err
		}
		driver.Age = uint32(age)
		drivers = append(drivers, driver)
	}
	return drivers, nil
}

func parseDrivesCSV(cr *csv.Reader) ([]model.Driver, error) {

	records, err := cr.ReadAll()
	if err != nil {
		return []model.Driver{}, err
	}
	drivers := []model.Driver{}
	for i, record := range records {
		if i == 0 {
			continue
		}
		fmt.Printf("FirstName: %s LastName %s\n", record[1], record[2])

		driver := model.Driver{}
		// Parse DriverID
		driverID, err := strconv.Atoi(record[0])
		if err != nil {
			panic(err)
		}
		driver.ID = uint64(driverID)
		driver.FirstName = record[1]
		driver.LastName = record[2]
		driver.Surname = record[3]
		// Parse date
		birthdate, err := time.Parse(birthdateLayout, record[4])

		if err != nil {
			fmt.Println(err)
		}
		driver.BirthDate = &birthdate

		driver.Gender = record[5]
		driver.Email = record[6]
		driver.PhoneNumber = record[7]
		driver.SideNumber = record[8]
		driver.EmergencyContact = record[9]
		driver.EmergencyNumber = record[10]
		// parse age
		age, err := strconv.Atoi(record[11])
		if err != nil {
			panic(err)
		}
		driver.Age = uint32(age)
		drivers = append(drivers, driver)

	}

	return drivers, nil
}
