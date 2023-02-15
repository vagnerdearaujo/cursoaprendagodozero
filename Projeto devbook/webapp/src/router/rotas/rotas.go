package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

type WebRotas struct {
	URI                string
	Metodo             string
	Funcao             func(w http.ResponseWriter, r *http.Request)
	RequerAutenticacao bool
}

func ConfigurarWebRotas(router *mux.Router) *mux.Router {
	rotas := rotaslogin
	
	for _, rota := range rotas {
		router.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return router

}
