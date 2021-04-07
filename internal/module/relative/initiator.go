package relative

import (
	"rideBenefit/internal/constant/model"
	"rideBenefit/internal/repository"
	"rideBenefit/internal/storage/persistence"
)

// Usecase contains the function of business logic of domain relative
type Usecase interface {
	GetRelative(relativeID uint64) (*model.Relative, error)
	AddRelative(relative *model.Relative) (*model.Relative, error)
	UpdateRelative(relative *model.Relative) (*model.Relative, error)
	DeleteRelative(relativeID uint64) error
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
