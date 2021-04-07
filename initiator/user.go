package initiator

import (
	routing "rideBenefit/internal/api/rest"
	"rideBenefit/internal/handler/rest"
	"rideBenefit/internal/module/user"
	"rideBenefit/internal/repository"
	"rideBenefit/internal/storage/persistence"
	"rideBenefit/platform/cockroach"
	"rideBenefit/platform/httprouter"
)

// User initializes the domain user
func User(cockroachPlatform cockroach.CockroachPlatform) []httprouter.Router {

	// Initiate the user persistence
	userPersistence := persistence.UserInit(cockroachPlatform)
	// Initiate the user repository
	userRepository := repository.UserInit()
	// Initiate the user service
	usecase := user.Initialize(userRepository, userPersistence)

	// Initiate the user rest API handler
	handler := rest.UserInit(usecase)

	// Initiate the user routing
	return routing.UserRouting(handler)

}
