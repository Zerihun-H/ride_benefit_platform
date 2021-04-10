package initiator

import (
	routing "rideBenefit/internal/api/rest"
	"rideBenefit/internal/handler/rest"
	"rideBenefit/internal/module/auth"
	"rideBenefit/internal/repository"
	"rideBenefit/internal/storage/persistence"
	"rideBenefit/platform/cockroach"
	"rideBenefit/platform/httprouter"
)

// Auth initializes the domain auth
func Auth(cockroachPlatform cockroach.CockroachPlatform) []httprouter.Router {

	// Initiate the auth persistence
	authPersistence := persistence.AuthInit(cockroachPlatform)
	// Initiate the auth repository
	authRepository := repository.AuthInit()
	// Initiate the auth service
	usecase := auth.Initialize(authRepository, authPersistence)

	// Initiate the auth rest API handler
	handler := rest.AuthInit(usecase)

	// Initiate the auth routing
	return routing.AuthRouting(handler)

}
