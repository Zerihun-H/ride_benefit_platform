package initiator

import (
	routing "rideBenefit/internal/api/rest"
	"rideBenefit/internal/handler/rest"
	"rideBenefit/internal/module/partner"
	"rideBenefit/internal/repository"
	"rideBenefit/internal/storage/persistence"
	"rideBenefit/platform/cockroach"
	"rideBenefit/platform/httprouter"
)

// Partner initializes the domain partner
func Partner(cockroachPlatform cockroach.CockroachPlatform) []httprouter.Router {

	// Initiate the partner persistence
	partnerPersistence := persistence.PartnerInit(cockroachPlatform)
	// Initiate the partner repository
	partnerRepository := repository.PartnerInit()
	// Initiate the partner service
	usecase := partner.Initialize(partnerRepository, partnerPersistence)

	// Initiate the partner rest API handler
	handler := rest.PartnerInit(usecase)

	// Initiate the partner routing
	return routing.PartnerRouting(handler)

}
