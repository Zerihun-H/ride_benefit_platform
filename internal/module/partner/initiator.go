package partner

import (
	"rideBenefit/internal/repository"
	"rideBenefit/internal/storage/persistence"
)

// Usecase contains the function of business logic of domain partner
type Usecase interface {
}

type service struct {
	partnerRepo    repository.PartnerRepository
	partnerPersist persistence.PartnerPersistence
}

// Initialize takes all necessary service for domain partner to run the business logic of domain partner
func Initialize(partnerRepo repository.PartnerRepository, partnerPersist persistence.PartnerPersistence) Usecase {
	return &service{
		partnerRepo:    partnerRepo,
		partnerPersist: partnerPersist,
	}
}
