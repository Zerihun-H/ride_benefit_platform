package user

import (
	"rideBenefit/internal/repository"
	"rideBenefit/internal/storage/persistence"
)

// Usecase contains the function of business logic of domain user
type Usecase interface {
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
