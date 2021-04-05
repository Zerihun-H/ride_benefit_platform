package initiator

import (
	"log"

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

	// Get the platform connection closer
	db := cockroachPlatform.Open()
	dbc, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbc.Close()

	// Initiate the partner persistence
	partnerPersistence := persistence.PartnerInit(db)
	// Initiate the partner repository
	partnerRepository := repository.PartnerInit()
	// Initiate the partner service
	usecase := partner.Initialize(partnerRepository, partnerPersistence)

	// Initiate the partner rest API handler
	handler := rest.PartnerInit(usecase)

	// Initiate the partner routing
	return routing.PartnerRouting(handler)

}
