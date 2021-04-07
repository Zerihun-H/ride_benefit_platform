package persistence

import (
	"rideBenefit/internal/constant/model"
	"rideBenefit/platform/cockroach"
)

// AuthPersistence contains the list of functions for database table auths
type AuthPersistence interface {
	GetRole(roleID uint64) (*model.Role, error)
	AddRole(role *model.Role) (*model.Role, error)
	GetUserByUsername(username string) (*model.User, error)
}

type authPersistence struct {
	db cockroach.CockroachPlatform
}

// AuthInit is to init the auth persistence that contains auth data
func AuthInit(db cockroach.CockroachPlatform) AuthPersistence {
	return &authPersistence{
		db,
	}
}

// GetRole using the auth id fetchs the role from the auth database
func (ap *authPersistence) GetRole(roleID uint64) (*model.Role, error) {
	db, err := ap.db.Open()
	if err != nil {
		return nil, err
	}
	dbc, err := db.DB()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()

	role := &model.Role{}
	if err := db.Where("id = ?", roleID).First(role).Error; err != nil {

		return &model.Role{}, err
	}
	return role, nil
}

// AddRole is adds a role to the database given a valid role
func (ap *authPersistence) AddRole(role *model.Role) (*model.Role, error) {

	db, err := ap.db.Open()
	if err != nil {
		return nil, err
	}
	dbc, err := db.DB()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()

	if err := db.Create(role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (ap *authPersistence) GetUserByUsername(username string) (*model.User, error) {
	db, err := ap.db.Open()
	if err != nil {
		return nil, err
	}
	dbc, err := db.DB()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()

	user := &model.User{}
	if err := db.Where("user_name = ?", username).First(user).Error; err != nil {

		return &model.User{}, err
	}

	return user, nil
}
