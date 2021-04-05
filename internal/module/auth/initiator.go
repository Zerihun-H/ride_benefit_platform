package auth

import (
	"rideBenefit/internal/repository"
	"rideBenefit/internal/storage/persistence"
)

// Usecase contains the function of business logic of domain auth
type Usecase interface {
}

type service struct {
	authRepo    repository.AuthRepository
	authPersist persistence.AuthPersistence
}

// Initialize takes all necessary service for domain auth to run the business logic of domain auth
func Initialize(authRepo repository.AuthRepository, authPersist persistence.AuthPersistence) Usecase {
	return &service{
		authRepo:    authRepo,
		authPersist: authPersist,
	}
}
