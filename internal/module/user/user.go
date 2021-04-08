package user

import (
	"rideBenefit/internal/constant/model"

	"golang.org/x/crypto/bcrypt"
)

func (s *service) GetUser(userID uint64) (*model.User, error) {
	// Some validation

	user, err := s.userPersist.GetUser(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) AddUser(user *model.User) (*model.User, error) {
	// Hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(passwordHash)
	usr, err := s.userPersist.AddUser(user)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (s *service) UpdateUser(user *model.User) (*model.User, error) {

	usr, err := s.userPersist.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (s *service) DeleteUser(userID uint64) error {

	return s.userPersist.DeleteUser(userID)
}
