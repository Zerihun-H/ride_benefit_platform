package driver

import (
	"rideBenefit/internal/repository"
	"rideBenefit/internal/storage/persistence"
)

// Usecase contains the function of business logic of domain driver
type Usecase interface {
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
