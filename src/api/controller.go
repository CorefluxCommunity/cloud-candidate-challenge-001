package api

import (
	"io"
	"net/http"
)

type Controller struct {
}

func (c Controller) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", c.HealthCheck)
	mux.HandleFunc("GET /create", c.CreateHandler)
	mux.HandleFunc("GET /destroy", c.DestroyHandler)
}
func (c *Controller) HealthCheck(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Health check ok")
}

func (c *Controller) CreateHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello from Create")
	// Recebe as configuraçoes para criar uma instancia no digital ocean em json e chama o serviço que se comunica com o terraform

	// responde o output do terraform para o cliente
}

func (c *Controller) DestroyHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello from Destroy")
	// recebe o ID do serviço que deseja remover

	// responde com o output do terraform
}
