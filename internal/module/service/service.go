package service

import "rideBenefit/internal/constant/model"

func (s *service) GetService(serviceID uint64) (*model.Service, error) {

	service, err := s.servicePersist.GetService(serviceID)
	if err != nil {
		return nil, err
	}

	return service, nil
}

func (s *service) AddService(service *model.Service) (*model.Service, error) {

	drv, err := s.servicePersist.AddService(service)
	if err != nil {
		return nil, err
	}

	return drv, nil
}

func (s *service) UpdateService(service *model.Service) (*model.Service, error) {

	part, err := s.servicePersist.UpdateService(service)
	if err != nil {
		return nil, err
	}

	return part, nil
}

func (s *service) DeleteService(serviceID uint64) error {

	return s.servicePersist.DeleteService(serviceID)
}
