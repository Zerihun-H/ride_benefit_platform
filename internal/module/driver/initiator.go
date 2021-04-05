package driver

import (
	"rideBenefit/internal/constant/model"
	"rideBenefit/internal/repository"
	"rideBenefit/internal/storage/persistence"
)

// Usecase contains the function of business logic of domain driver
type Usecase interface {
	GetDriver(driverID uint64) (*model.Driver, error)
	AddDriver(driver *model.Driver) (*model.Driver, error)
	UpdateDriver(driver *model.Driver) (*model.Driver, error)
	DeleteDriver(driverID uint64) error
	AddDrivers(drivers []model.Driver) error
}

type service struct {
	driverRepo    repository.DriverRepository
	driverPersist persistence.DriverPersistence
}

// Initialize takes all necessary service for domain driver to run the business logic of domain driver
func Initialize(driverRepo repository.DriverRepository, driverPersist persistence.DriverPersistence) Usecase {
	return &service{
		driverRepo:    driverRepo,
		driverPersist: driverPersist,
	}
}
