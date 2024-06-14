package main

import (
	"log"

	"github.com/Desgue/cloud-candidate-challenge-001/src/api"
)

func main() {
	server := api.NewServer(":8000")
	log.Fatal(server.Start())
}
