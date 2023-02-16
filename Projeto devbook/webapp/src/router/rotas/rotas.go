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
	rotas := rotasLogin
	rotas = append(rotas, rotasUsuarios...)

	for _, rota := range rotas {
		router.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	//Configuração do File Server.
	//Esta configuração tem por objetivo auxiliar na localização de arquivos de recursos (assets)
	//=> Note que o único lugar onde o "." vai aparecer é: http.Dir("./assets/")

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router

}
