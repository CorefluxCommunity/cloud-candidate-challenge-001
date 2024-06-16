package main

import (
	"log"

	"github.com/Desgue/cloud-candidate-challenge-001/src/api"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := api.NewServer(":8000", api.NewDropletController())
	log.Fatal(server.Start())
}
