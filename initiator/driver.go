package initiator

import (
	"log"

	routing "rideBenefit/internal/api/rest"
	"rideBenefit/internal/handler/rest"
	"rideBenefit/internal/module/driver"
	"rideBenefit/internal/repository"
	"rideBenefit/internal/storage/persistence"
	"rideBenefit/platform/cockroach"
	"rideBenefit/platform/httprouter"
)

// Driver initializes the domain driver
func Driver(cockroachPlatform cockroach.CockroachPlatform) []httprouter.Router {

	// Get the platform closer
	db := cockroachPlatform.Open()
	dbc, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbc.Close()

	// Initiate the driver persistence
	driverPersistence := persistence.DriverInit(db)
	// Initiate the driver repository
	driverRepository := repository.DriverInit()
	// Initiate the driver service
	usecase := driver.Initialize(driverRepository, driverPersistence)

	// Initiate the driver rest API handler
	handler := rest.DriverInit(usecase)

	// Initiate the driver routing
	return routing.DriverRouting(handler)

}
