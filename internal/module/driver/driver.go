package driver

import "rideBenefit/internal/constant/model"

func (s *service) GetDriver(driverID uint64) (*model.Driver, error) {
	// Some validation

	driver, err := s.driverPersist.GetDriver(driverID)
	if err != nil {
		return nil, err
	}

	return driver, nil
}

func (s *service) AddDriver(driver *model.Driver) (*model.Driver, error) {
	// Check if the drivered pruduct is valid and available
	// Check if the customer is valid
	// Check if the shop is valid
	drv, err := s.driverPersist.AddDriver(driver)
	if err != nil {
		return nil, err
	}

	return drv, nil
}
