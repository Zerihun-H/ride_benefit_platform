package persistence

import (
	"rideBenefit/internal/constant/model"

	"gorm.io/gorm"
)

// ServicePersistence contains the list of functions for database table services
type ServicePersistence interface {
	GetService(serviceID uint64) (*model.Service, error)
	AddService(service *model.Service) (*model.Service, error)
}

type servicePersistence struct {
	db *gorm.DB
}

// ServiceInit is to init the service persistence that contains service data
func ServiceInit(db *gorm.DB) ServicePersistence {
	return &servicePersistence{
		db,
	}
}

// Getservice using the service id fetchs the service from the service database
func (pp *servicePersistence) GetService(serviceID uint64) (*model.Service, error) {
	service := &model.Service{}
	if err := pp.db.Where("id = ?", serviceID).First(service).Error; err != nil {

		return &model.Service{}, err
	}
	return service, nil
}

// AddService is adds a service to the database given a valid diver
func (pp *servicePersistence) AddService(service *model.Service) (*model.Service, error) {
	if err := pp.db.Create(service).Error; err != nil {
		return nil, err
	}
	return service, nil
}
