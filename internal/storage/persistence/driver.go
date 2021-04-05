package persistence

import (
	"rideBenefit/internal/constant/model"

	"gorm.io/gorm"
)

// DriverPersistence contains the list of functions for database table drivers
type DriverPersistence interface {
	GetDriver(driverID uint64) (*model.Driver, error)
	AddDriver(driver *model.Driver) (*model.Driver, error)
	UpdateDriver(driver *model.Driver) (*model.Driver, error)
	DeleteDriver(driverID uint64) error
	AddDrivers(drivers []model.Driver) error
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

// AddDriver is adds a driver to the database given a valid diver
func (dp *driverPersistence) UpdateDriver(driver *model.Driver) (*model.Driver, error) {
	updatedDriver := driver
	err := dp.db.First(driver).Error
	if err != nil {
		return nil, err
	}
	updatedDriver.ID = driver.ID
	dp.db.Save(&updatedDriver)
	return driver, nil
}

// DeleteDriver is adds a driver to the database given a valid diver
func (dp *driverPersistence) DeleteDriver(driverID uint64) error {

	return dp.db.Where("id = ?", driverID).Delete(&model.Driver{}).Error
}

// AddDrivers is adds a drivers to the database given valid divers
func (dp *driverPersistence) AddDrivers(drivers []model.Driver) error {
	return dp.db.Create(drivers).Error
}
