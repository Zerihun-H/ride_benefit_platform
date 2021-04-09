package persistence

import (
	"fmt"
	"rideBenefit/internal/constant/model"
	"rideBenefit/platform/cockroach"
	"strconv"
)

// AuthPersistence contains the list of functions for database table auths
type AuthPersistence interface {
	GetRole(roleID uint64) (*model.Role, error)
	AddRole(role *model.Role) (*model.Role, error)
	GetUserByUsername(username string) (*model.User, error)
	GetRolePermissions(roleID uint64) ([]model.Permission, error)
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
	if err := db.Where("username = ?", username).First(user).Error; err != nil {

		return &model.User{}, err
	}

	return user, nil
}

func (ap *authPersistence) GetRolePermissions(roleID uint64) ([]model.Permission, error) {
	db, err := ap.db.Open()
	if err != nil {
		return nil, err
	}
	dbc, err := db.DB()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()
	permissions := []model.Permission{}
	// Parse RoleID to string
	rid := strconv.Itoa(int(roleID))
	// Prerpare query statement
	qry := fmt.Sprintf("SELECT name,description FROM public.role_permissions LEFT JOIN public.permissions ON public.role_permissions.permission_id = public.permissions.id where public.role_permissions.role_id = %v;", rid)
	err = db.Raw(qry).Scan(&permissions).Error
	if err != nil {
		return nil, err
	}

	return permissions, nil
}
