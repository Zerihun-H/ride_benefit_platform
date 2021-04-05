package persistence

import (
	"rideBenefit/internal/constant/model"

	"gorm.io/gorm"
)

// UserPersistence contains the list of functions for database table users
type UserPersistence interface {
	GetUser(userID uint64) (*model.User, error)
	AddUser(user *model.User) (*model.User, error)
}

type userPersistence struct {
	db *gorm.DB
}

// UserInit is to init the user persistence that contains user data
func UserInit(db *gorm.DB) UserPersistence {
	return &userPersistence{
		db,
	}
}

// Getuser using the user id fetchs the user from the user database
func (pp *userPersistence) GetUser(userID uint64) (*model.User, error) {
	user := &model.User{}
	if err := pp.db.Where("id = ?", userID).First(user).Error; err != nil {

		return &model.User{}, err
	}
	return user, nil
}

// AddUser is adds a user to the database given a valid diver
func (pp *userPersistence) AddUser(user *model.User) (*model.User, error) {
	if err := pp.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
