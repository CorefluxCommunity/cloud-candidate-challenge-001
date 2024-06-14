package api

import "net/http"

func createHandler(w http.ResponseWriter, r *http.Request) {
	// Recebe as configuraçoes para criar uma instancia no digital ocean em json e chama o serviço que se comunica com o terraform

	// responde o output do terraform para o cliente
}

func destroyHandler(w http.ResponseWriter, r *http.Request) {
	// recebe o ID do serviço que deseja remover

	// responde com o output do terraform
}
