package create_droplet

import (
	"fmt"
	"net/http"
	"src/services"
	"sync"
)

func CreateDropletHandler(w http.ResponseWriter, r *http.Request, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	go func() {
		token, err := services.GetDigitalOceanToken()
		if err != nil {
			ch <- fmt.Sprintf("Error: %v", err)
			return
		}

		req, err := services.DecodeRequest(r)
		if err != nil {
			ch <- fmt.Sprintf("Error: %v", err)
			return
		}

		output, err := services.RunTerraformApply(req, token)
		if err != nil {
			ch <- fmt.Sprintf("Error: %v, Output: %s", err, string(output))
			return
		}

		ipAddress, err := services.GetDropletIPAddress()
		if err != nil {
			ch <- fmt.Sprintf("Error: %v", err)
			return
		}

		ch <- ipAddress
	}()
}

func CreateDropletRoutine(w http.ResponseWriter, r *http.Request) {
	ch := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	go CreateDropletHandler(w, r, ch, &wg)
	wg.Wait()
	services.HandleHTTPResponse(w, ch)
	close(ch)
}
