package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota Modelo para a criação das rotas
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// ConfigurarRotas inclui todas as rotas configuradas
func ConfigurarRotas(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotaLogin)

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}
