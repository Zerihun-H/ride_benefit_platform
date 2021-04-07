package initiator

import (
	"fmt"
	"log"
	"os"
	"rideBenefit/platform/cockroach"
	"rideBenefit/platform/httprouter"
)

func CockroachInitiator() cockroach.CockroachPlatform {
	// Initaite cockroch platform
	cockroachPlatform := cockroach.Initialize(dbURL)
	// Migrate necessary tables
	err := cockroachPlatform.Migrate()
	if err != nil {
		log.Fatal(err)
	}
	return cockroachPlatform
}

func Initiator() {
	// Initiate cockroach
	cockroahPlatform := CockroachInitiator()

	// Initiate employee module and
	employeeRouters := Employee(cockroahPlatform)
	partnerRouters := Partner(cockroahPlatform)

	routers := []httprouter.Router{}
	routers = append(routers, employeeRouters...)
	routers = append(routers, partnerRouters...)

	// Get self host port
	hostPort := os.Getenv("SELF_PORT")
	hostAddress := os.Getenv("SELF_ADDRESS")
	hostURL := fmt.Sprintf(hostAddress + ":" + hostPort)

	// Get the allowed request origins for the http server
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	// Initiate the http server
	server := httprouter.Initialize(hostURL, allowedOrigins, routers, domain)

	// Get the handlers from

	// Start the http server
	server.Serve()
}
