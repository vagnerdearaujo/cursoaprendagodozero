package router

import (
	"webapp/src/router/rotas"

	"github.com/gorilla/mux"
)

func GerarRotasWebAPP() *mux.Router {
	r := mux.NewRouter()
	return rotas.ConfigurarWebRotas(r)
}
