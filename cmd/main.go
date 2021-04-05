package main

import (
	"rideBenefit/initiator"

	_ "github.com/lib/pq"
)

func main() {
	initiator.Initiator()
}
