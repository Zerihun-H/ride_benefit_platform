package driver

import (
	"rideBenefit/internal/constant/model"
)

func (s *service) GetDriver(driverID uint64) (*model.Driver, error) {
	// Some validation

	return s.driverPersist.GetDriver(driverID)
}

func (s *service) AddDriver(driver *model.Driver) (*model.Driver, error) {

	return s.driverPersist.AddDriver(driver)
}

func (s *service) UpdateDriver(driver *model.Driver) (*model.Driver, error) {

	return s.driverPersist.UpdateDriver(driver)

}

func (s *service) DeleteDriver(driverID uint64) error {

	return s.driverPersist.DeleteDriver(driverID)

}

func (s *service) AddDrivers(drivers []model.Driver) error {

	return s.driverPersist.AddDrivers(drivers)
}
