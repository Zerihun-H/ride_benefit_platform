package relative

import (
	"rideBenefit/internal/repository"
	"rideBenefit/internal/storage/persistence"
)

// Usecase contains the function of business logic of domain relative
type Usecase interface {
}

type service struct {
	relativeRepo    repository.RelativeRepository
	relativePersist persistence.RelativePersistence
}

// Initialize takes all necessary service for domain relative to run the business logic of domain relative
func Initialize(relativeRepo repository.RelativeRepository, relativePersist persistence.RelativePersistence) Usecase {
	return &service{
		relativeRepo:    relativeRepo,
		relativePersist: relativePersist,
	}
}
