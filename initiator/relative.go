package initiator

import (
	routing "rideBenefit/internal/api/rest"
	"rideBenefit/internal/handler/rest"
	"rideBenefit/internal/module/relative"
	"rideBenefit/internal/repository"
	"rideBenefit/internal/storage/persistence"
	"rideBenefit/platform/cockroach"
	"rideBenefit/platform/httprouter"
)

// Relative initializes the domain relative
func Relative(cockroachPlatform cockroach.CockroachPlatform) []httprouter.Router {

	// Initiate the relative persistence
	relativePersistence := persistence.RelativeInit(cockroachPlatform)
	// Initiate the relative repository
	relativeRepository := repository.RelativeInit()
	// Initiate the relative service
	usecase := relative.Initialize(relativeRepository, relativePersistence)

	// Initiate the relative rest API handler
	handler := rest.RelativeInit(usecase)

	// Initiate the relative routing
	return routing.RelativeRouting(handler)

}
