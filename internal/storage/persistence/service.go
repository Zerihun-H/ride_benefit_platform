package persistence

import (
	"rideBenefit/internal/constant/model"
	"rideBenefit/platform/cockroach"
)

// ServicePersistence contains the list of functions for database table services
type ServicePersistence interface {
	GetService(serviceID uint64) (*model.Service, error)
	AddService(service *model.Service) (*model.Service, error)
	UpdateService(service *model.Service) (*model.Service, error)
	DeleteService(serviceID uint64) error
}

type servicePersistence struct {
	db cockroach.CockroachPlatform
}

// ServiceInit is to init the service persistence that contains service data
func ServiceInit(db cockroach.CockroachPlatform) ServicePersistence {
	return &servicePersistence{
		db,
	}
}

// Getservice using the service id fetchs the service from the service database
func (pp *servicePersistence) GetService(serviceID uint64) (*model.Service, error) {
	db, err := pp.db.Open()
	if err != nil {
		return nil, err
	}
	dbc, err := db.DB()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()
	service := &model.Service{}
	if err := db.Where("id = ?", serviceID).First(service).Error; err != nil {

		return &model.Service{}, err
	}

	return service, nil
}

// AddService is adds a service to the database given a valid diver
func (pp *servicePersistence) AddService(service *model.Service) (*model.Service, error) {
	db, err := pp.db.Open()
	if err != nil {
		return nil, err
	}
	dbc, err := db.DB()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()
	if err := db.Create(service).Error; err != nil {
		return nil, err
	}

	return service, nil
}

// AddService is adds a service to the database given a valid diver
func (sp *servicePersistence) UpdateService(service *model.Service) (*model.Service, error) {
	db, err := sp.db.Open()
	if err != nil {
		return nil, err
	}
	dbc, err := db.DB()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()
	updatedService := *service
	err = db.First(service).Error
	if err != nil {
		return nil, err
	}

	updatedService.ID = service.ID
	err = db.Save(&updatedService).Error
	if err != nil {
		return nil, err
	}
	return &updatedService, nil
}

// AddService is adds a service to the database given a valid diver
func (sp *servicePersistence) DeleteService(serviceID uint64) error {
	db, err := sp.db.Open()
	if err != nil {
		return err
	}
	dbc, err := db.DB()
	if err != nil {
		return err
	}
	defer dbc.Close()

	return db.Where("id = ?", serviceID).Delete(&model.Service{}).Error
}
