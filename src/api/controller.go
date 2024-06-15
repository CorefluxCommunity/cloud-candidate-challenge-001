package api

import (
	"io"
	"net/http"
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
	io.WriteString(w, "Hello from Create")
	// Recebe as configuraçoes para criar uma instancia no digital ocean em json e chama o serviço que se comunica com o terraform

	// responde o output do terraform para o cliente
}

func (c *Controller) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello from Destroy")
	// recebe o ID do serviço que deseja remover

	// responde com o output do terraform
}
