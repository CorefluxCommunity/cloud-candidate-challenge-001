package main

import (
	"awesomeProject/functions"
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		return
	}
	fmt.Println("Server running...")
	http.HandleFunc("/create", functions.CreateDropletHandler)
	http.HandleFunc("/all", functions.ListDropletsHandler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
