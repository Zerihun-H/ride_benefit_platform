package persistence

import (
	"rideBenefit/internal/constant/model"
	"rideBenefit/platform/cockroach"
)

// UserPersistence contains the list of functions for database table users
type UserPersistence interface {
	GetUser(userID uint64) (*model.User, error)
	AddUser(user *model.User) (*model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	DeleteUser(userID uint64) error
}

type userPersistence struct {
	db cockroach.CockroachPlatform
}

// UserInit is to init the user persistence that contains user data
func UserInit(db cockroach.CockroachPlatform) UserPersistence {
	return &userPersistence{
		db,
	}
}

// Getuser using the user id fetchs the user from the user database
func (up *userPersistence) GetUser(userID uint64) (*model.User, error) {
	db, err := up.db.Open()
	if err != nil {
		return nil, err
	}
	dbc, err := db.DB()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()
	user := &model.User{}
	if err := db.Where("id = ?", userID).First(user).Error; err != nil {

		return &model.User{}, err
	}

	return user, nil
}

// AddUser is adds a user to the database given a valid diver
func (up *userPersistence) AddUser(user *model.User) (*model.User, error) {
	db, err := up.db.Open()
	if err != nil {
		return nil, err
	}
	dbc, err := db.DB()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()
	if err := db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// AddUser is adds a user to the database given a valid diver
func (up *userPersistence) UpdateUser(user *model.User) (*model.User, error) {
	db, err := up.db.Open()
	if err != nil {
		return nil, err
	}
	dbc, err := db.DB()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()
	updatedUser := *user
	err = db.First(user).Error
	if err != nil {
		return nil, err
	}

	updatedUser.ID = user.ID
	err = db.Save(&updatedUser).Error
	if err != nil {
		return nil, err
	}
	return &updatedUser, nil
}

// AddUser is adds a user to the database given a valid diver
func (up *userPersistence) DeleteUser(userID uint64) error {
	db, err := up.db.Open()
	if err != nil {
		return err
	}
	dbc, err := db.DB()
	if err != nil {
		return err
	}
	defer dbc.Close()

	return db.Where("id = ?", userID).Delete(&model.User{}).Error
}
