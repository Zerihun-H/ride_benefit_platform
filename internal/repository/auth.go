package repository

import "rideBenefit/internal/constant/model"

var RolePermssions model.RolesPermissions

// AuthRepository contains the functions of data logic for domain auth
type AuthRepository interface {
	GetRolesPermissions() (*model.RolesPermissions, error)
	LoadRolesPermisions() (*model.RolesPermissions, error)
}

type authRepository struct {
}

// AuthInit initializes the data logic / repository for domain auth
func AuthInit() AuthRepository {

	return &authRepository{}
}

func (ar *authRepository) GetRolesPermissions() (*model.RolesPermissions, error) {
	return nil, nil
}

func (ar *authRepository) LoadRolesPermisions() (*model.RolesPermissions, error) {

	return nil, nil
}


