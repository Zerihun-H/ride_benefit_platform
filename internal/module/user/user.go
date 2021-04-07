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

	part, err := s.userPersist.AddUser(user)
	if err != nil {
		return nil, err
	}

	return part, nil
}

func (s *service) UpdateUser(user *model.User) (*model.User, error) {

	part, err := s.userPersist.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return part, nil
}

func (s *service) DeleteUser(userID uint64) error {

	return s.userPersist.DeleteUser(userID)
}
