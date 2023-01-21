package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// GerarRotas retona um router com as rotas configuradas.
func GerarRotas() *mux.Router {
	r := mux.NewRouter()
	return rotas.ConfigurarRotas(r)
}
