package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Desgue/cloud-candidate-challenge-001/src/droplet"
	"github.com/Desgue/cloud-candidate-challenge-001/src/svc"
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
	mux.HandleFunc("GET /droplet", c.GetHandler)
	mux.HandleFunc("POST /droplet", c.PostHandler)
	mux.HandleFunc("DELETE /droplet", c.DeleteHandler)
}
func (c *DropletController) GetHandler(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Health check ok")
}

func (c *DropletController) PostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Recebe as configuraçoes para criar uma instancia no digital ocean em json e chama o serviço que se comunica com o terraform
	var dropletReq droplet.DropletRequest
	err := json.NewDecoder(r.Body).Decode(&dropletReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dropletRes, err := c.DropletService.CreateDroplet(dropletReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// responde o output do terraform para o cliente
	err = json.NewEncoder(w).Encode(dropletRes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (c *DropletController) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// recebe o ID do serviço que deseja remover
	err := c.DropletService.DeleteDroplet()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write([]byte("Droplet Deleted sucessfully"))

}
