package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Desgue/cloud-candidate-challenge-001/src/terraform"
)

type Controller struct {
}

func (c Controller) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /droplet", c.HealthCheck)
	mux.HandleFunc("POST /droplet", c.PostHandler)
	mux.HandleFunc("DELETE /droplet", c.DeleteHandler)
}
func (c *Controller) HealthCheck(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Health check ok")
}

func (c *Controller) PostHandler(w http.ResponseWriter, r *http.Request) {
	// Recebe as configuraçoes para criar uma instancia no digital ocean em json e chama o serviço que se comunica com o terraform
	var dropletReq terraform.DropletRequest
	err := json.NewDecoder(r.Body).Decode(&dropletReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !dropletReq.IsValid() {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	// responde o output do terraform para o cliente
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dropletReq)

}

func (c *Controller) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello from Destroy")
	// recebe o ID do serviço que deseja remover

	// responde com o output do terraform
}
