package activity

import (
	"rideBenefit/internal/repository"
	"rideBenefit/internal/storage/persistence"
)

// Usecase contains the function of business logic of domain activity
type Usecase interface {
}

type service struct {
	activityRepo    repository.ActivityRepository
	activityPersist persistence.ActivityPersistence
}

// Initialize takes all necessary service for domain activity to run the business logic of domain activity
func Initialize(activityRepo repository.ActivityRepository, activityPersist persistence.ActivityPersistence) Usecase {
	return &service{
		activityRepo:    activityRepo,
		activityPersist: activityPersist,
	}
}
