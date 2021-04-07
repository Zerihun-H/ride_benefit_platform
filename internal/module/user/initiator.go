package user

import (
	"rideBenefit/internal/constant/model"
	"rideBenefit/internal/repository"
	"rideBenefit/internal/storage/persistence"
)

// Usecase contains the function of business logic of domain user
type Usecase interface {
	GetUser(userID uint64) (*model.User, error)
	AddUser(user *model.User) (*model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	DeleteUser(userID uint64) error
}

type service struct {
	userRepo    repository.UserRepository
	userPersist persistence.UserPersistence
}

// Initialize takes all necessary user for domain user to run the business logic of domain user
func Initialize(userRepo repository.UserRepository, userPersist persistence.UserPersistence) Usecase {
	return &service{
		userRepo:    userRepo,
		userPersist: userPersist,
	}
}
