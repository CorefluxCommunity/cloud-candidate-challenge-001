package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Desgue/cloud-candidate-challenge-001/src/svc"
	"github.com/Desgue/cloud-candidate-challenge-001/src/terraform"
)

type DropletController struct {
	DropletService *svc.DropletService
}

func NewDropletController() *DropletController {
	return &DropletController{
		DropletService: svc.NewDropletService(),
	}
}

func (c DropletController) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /droplet", c.HealthCheck)
	mux.HandleFunc("POST /droplet", c.PostHandler)
	mux.HandleFunc("DELETE /droplet", c.DeleteHandler)
}
func (c *DropletController) HealthCheck(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Health check ok")
}

func (c *DropletController) PostHandler(w http.ResponseWriter, r *http.Request) {
	// Recebe as configuraçoes para criar uma instancia no digital ocean em json e chama o serviço que se comunica com o terraform
	var dropletReq terraform.DropletRequest
	err := json.NewDecoder(r.Body).Decode(&dropletReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// responde o output do terraform para o cliente
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dropletReq)

}

func (c *DropletController) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello from Destroy")
	// recebe o ID do serviço que deseja remover

	// responde com o output do terraform
}
