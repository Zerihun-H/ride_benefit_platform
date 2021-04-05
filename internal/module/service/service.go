package service

import "rideBenefit/internal/constant/model"

func (s *service) GetService(serviceID uint64) (*model.Service, error) {
	// Some validation

	service, err := s.servicePersist.GetService(serviceID)
	if err != nil {
		return nil, err
	}

	return service, nil
}

func (s *service) AddService(service *model.Service) (*model.Service, error) {
	// Check if the serviceed pruduct is valid and available
	// Check if the customer is valid
	// Check if the shop is valid
	drv, err := s.servicePersist.AddService(service)
	if err != nil {
		return nil, err
	}

	return drv, nil
}
