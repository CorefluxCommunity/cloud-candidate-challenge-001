package list_droplets

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"src/services"
	"sync"
)

func ListDropletsHandler(w http.ResponseWriter, r *http.Request, ch chan<- string, wg *sync.WaitGroup) {

	defer wg.Done()

	go func() {
		token := os.Getenv("DO_TOKEN")
		if token == "" {
			http.Error(w, "DigitalOcean token not found. Please set the DO_TOKEN environment variable.", http.StatusInternalServerError)
			return
		}

		output, err := services.RunTerraformListDroplets(token)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error: %v, Output: %s", err, string(output)), http.StatusInternalServerError)
			return
		}

		terraformOutput, err := services.GetTerraformOutput("terraform/list-droplets")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error getting output: %v, Output: %s", err, string(terraformOutput)), http.StatusInternalServerError)
			return
		}

		droplets, err := services.ParseDropletListOutput(terraformOutput)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error parsing output: %v", err), http.StatusInternalServerError)
			return
		}

		response, err := json.Marshal(droplets)
		if err != nil {
			ch <- fmt.Sprintf("Error encoding response: %v", err)
			return
		}

		ch <- string(response)
	}()
}

func ListDropletsRoutine(w http.ResponseWriter, r *http.Request) {
	ch := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	go ListDropletsHandler(w, r, ch, &wg)
	wg.Wait()
	services.HandleHTTPResponse(w, ch)
	close(ch)
}
