package auth

import "rideBenefit/internal/constant/model"

func (s *service) GetRoles(roleID uint64) (*model.Role, error) {

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

func (s *service) RolePermissions(roleID uint64) ([]model.Permission, error) {
	return s.authPersist.GetRolePermissions(roleID)
}

func (s *service) RoleHasPermission(roleID uint64, permission string) (bool, error) {

	// Get the role's permissions
	permissions, err := s.authPersist.GetRolePermissions(roleID)
	if err != nil {
		return false, err
	}

	// Check if the role has the requested permission
	for _, p := range permissions {
		if permission == p.Name {
			return true, nil
		}
	}

	return false, nil
}
