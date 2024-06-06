package main

import (
	"awesomeProject/functions"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server running...")
	http.HandleFunc("/create", functions.CreateDropletHandler)
	http.HandleFunc("/all", functions.ListDropletsHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
