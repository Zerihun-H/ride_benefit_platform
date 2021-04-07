package initiator

import (
	routing "rideBenefit/internal/api/rest"
	"rideBenefit/internal/handler/rest"
	"rideBenefit/internal/module/service"
	"rideBenefit/internal/repository"
	"rideBenefit/internal/storage/persistence"
	"rideBenefit/platform/cockroach"
	"rideBenefit/platform/httprouter"
)

// Service initializes the domain service
func Service(cockroachPlatform cockroach.CockroachPlatform) []httprouter.Router {

	// Initiate the service persistence
	servicePersistence := persistence.ServiceInit(cockroachPlatform)
	// Initiate the service repository
	serviceRepository := repository.ServiceInit()
	// Initiate the service service
	usecase := service.Initialize(serviceRepository, servicePersistence)

	// Initiate the service rest API handler
	handler := rest.ServiceInit(usecase)

	// Initiate the service routing
	return routing.ServiceRouting(handler)

}
