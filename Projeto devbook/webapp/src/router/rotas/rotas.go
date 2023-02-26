package rotas

import (
	"net/http"
	"webapp/src/middlewares"

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
	rotas = append(rotas, rotasPublicacoes...)

	for _, rota := range rotas {
		if rota.RequerAutenticacao {
			/*
				Chama Logger e passa ValidarInformacaoCookie como próxima função
				Chama ValidarInformacaoCookie e passa rota.Funcao como próxima função

				Assim: Logger chama ValidarInformacaoCookie que chama rota.Funcao.
			*/
			router.HandleFunc(rota.URI, middlewares.Logger(middlewares.ValidaInformacaoCookie(rota.Funcao))).Methods(rota.Metodo)
		} else {
			router.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
	}

	//Configuração do File Server.
	//Esta configuração tem por objetivo auxiliar na localização de arquivos de recursos (assets)
	//=> Note que o único lugar onde o "." vai aparecer é: http.Dir("./assets/")

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router

}
