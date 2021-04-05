package service

import (
	"rideBenefit/internal/repository"
	"rideBenefit/internal/storage/persistence"
)

// Usecase contains the function of business logic of domain service
type Usecase interface {
}

type service struct {
	serviceRepo    repository.ServiceRepository
	servicePersist persistence.ServicePersistence
}

// Initialize takes all necessary service for domain service to run the business logic of domain service
func Initialize(serviceRepo repository.ServiceRepository, servicePersist persistence.ServicePersistence) Usecase {
	return &service{
		serviceRepo:    serviceRepo,
		servicePersist: servicePersist,
	}

}
