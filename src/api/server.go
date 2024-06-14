package api

import (
	"io"
	"log"
	"net/http"
)

type server struct {
	addr string
}

func NewServer(addr string) *server {
	return &server{
		addr: addr,
	}
}

func (s *server) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Pong")
	})
	server := http.Server{
		Addr:    s.addr,
		Handler: mux,
	}
	log.Println("Starting server on", s.addr)
	return server.ListenAndServe()
}
