package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"src/functions"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}

func main() {
	fmt.Println("Server running...")
	http.HandleFunc("/create", functions.CreateDropletHandler)
	http.HandleFunc("/all", functions.ListDropletsHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
