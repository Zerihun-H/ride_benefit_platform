package persistence

import (
	"rideBenefit/internal/constant/model"

	"gorm.io/gorm"
)

// AuthPersistence contains the list of functions for database table auths
type AuthPersistence interface {
	GetRole(roleID uint64) (*model.Role, error)
	AddRole(role *model.Role) (*model.Role, error)
}

type authPersistence struct {
	db *gorm.DB
}

// AuthInit is to init the auth persistence that contains auth data
func AuthInit(db *gorm.DB) AuthPersistence {
	return &authPersistence{
		db,
	}
}

// GetRole using the auth id fetchs the role from the auth database
func (pp *authPersistence) GetRole(authID uint64) (*model.Role, error) {
	auth := &model.Role{}
	if err := pp.db.Where("id = ?", authID).First(auth).Error; err != nil {

		return &model.Role{}, err
	}
	return auth, nil
}

// AddRole is adds a role to the database given a valid role
func (pp *authPersistence) AddRole(auth *model.Role) (*model.Role, error) {
	if err := pp.db.Create(auth).Error; err != nil {
		return nil, err
	}
	return auth, nil
}
