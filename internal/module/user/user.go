package user

import "rideBenefit/internal/constant/model"

func (s *service) GetUser(userID uint64) (*model.User, error) {
	// Some validation

	user, err := s.userPersist.GetUser(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) AddUser(user *model.User) (*model.User, error) {
	// Check if the usered pruduct is valid and available
	// Check if the customer is valid
	// Check if the shop is valid
	drv, err := s.userPersist.AddUser(user)
	if err != nil {
		return nil, err
	}

	return drv, nil
}
