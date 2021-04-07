package service

import (
	"rideBenefit/internal/constant/model"
	"rideBenefit/internal/repository"
	"rideBenefit/internal/storage/persistence"
)

// Usecase contains the function of business logic of domain service
type Usecase interface {
	GetService(serviceID uint64) (*model.Service, error)
	AddService(service *model.Service) (*model.Service, error)
	UpdateService(service *model.Service) (*model.Service, error)
	DeleteService(serviceID uint64) error
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
