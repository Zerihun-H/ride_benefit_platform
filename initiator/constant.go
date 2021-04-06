package initiator

import (
	"fmt"
	"os"
)

const (
	cockroachURL = "postgres://%v:%v@%v:%v/%v?sslmode=disable"
	domain       = "rideBenefit"
)

var (
	dbUser = os.Getenv("CR_USER")
	dbName = os.Getenv("CR_NAME")
	dbPass = os.Getenv("CR_PASS")
	dbHost = os.Getenv("CR_HOST")
	dbPort = os.Getenv("CR_PORT")
	dbURL  = fmt.Sprintf(cockroachURL, dbUser, dbPass, dbHost, dbPort, dbName)
)
