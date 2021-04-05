package persistence

import (
	"rideBenefit/internal/constant/model"

	"gorm.io/gorm"
)

// DriverPersistence contains the list of functions for database table drivers
type DriverPersistence interface {
	GetDriver(driverID uint64) (*model.Driver, error)
	AddDriver(driver *model.Driver) (*model.Driver, error)
}

type driverPersistence struct {
	db *gorm.DB
}

// DriverInit is to init the driver persistence that contains driver data
func DriverInit(db *gorm.DB) DriverPersistence {
	return &driverPersistence{
		db,
	}
}

// Getdriver using the driver id fetchs the driver from the driver database
func (dp *driverPersistence) GetDriver(driverID uint64) (*model.Driver, error) {
	driver := &model.Driver{}
	if err := dp.db.Where("id = ?", driverID).First(driver).Error; err != nil {

		return &model.Driver{}, err
	}
	return driver, nil
}

// AddDriver is adds a driver to the database given a valid diver
func (dp *driverPersistence) AddDriver(driver *model.Driver) (*model.Driver, error) {
	if err := dp.db.Create(driver).Error; err != nil {
		return nil, err
	}
	return driver, nil
}
