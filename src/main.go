package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"src/services/create-droplet"
	"src/services/list-droplets"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}

func main() {
	fmt.Println("Server running...")
	http.HandleFunc("/create", create_droplet.CreateDropletRoutine)
	http.HandleFunc("/all", list_droplets.ListDropletsRoutine)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
