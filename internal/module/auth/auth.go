package auth

import "rideBenefit/internal/constant/model"

func (s *service) GetRoles(roleID uint64) (*model.Role, error) {
	// Some validation

	auth, err := s.authPersist.GetRole(roleID)
	if err != nil {
		return nil, err
	}

	return auth, nil
}

func (s *service) AddRole(auth *model.Role) (*model.Role, error) {

	drv, err := s.authPersist.AddRole(auth)
	if err != nil {
		return nil, err
	}

	return drv, nil
}