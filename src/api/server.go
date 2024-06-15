package api

import (
	"log"
	"net/http"
)

type server struct {
	addr       string
	controller *DropletController
}

func NewServer(addr string, controller *DropletController) *server {
	return &server{
		addr:       addr,
		controller: controller,
	}
}

func (s *server) Start() error {
	mux := http.NewServeMux()
	s.controller.registerRoutes(mux)

	server := http.Server{
		Addr:    s.addr,
		Handler: mux,
	}
	log.Println("Starting server on", s.addr)
	return server.ListenAndServe()
}
