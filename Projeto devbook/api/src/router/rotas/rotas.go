package rotas

import (
	"api/src/middlewares"
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
	//Ao invés de executar um loop para incluir cada uma das rotas, pode se incluir um ... para que o append
	//entenda que se trata de uma adição de slice do mesmo tipo.
	rotas = append(rotas, RotaSeguidores...)
	rotas = append(rotas, RotaPublicacoes...)

	/*
		for _, rota := range RotaSeguidores {
			rotas = append(rotas, rota)
		}
		for _, rota := range RotaPublicacoes {
			rotas = append(rotas, rota)
		}
	*/

	//Se houvesse uma função de logger que devesse ser chamada antes da autenticação:
	//r.HandleFunc(rota.URI, middlewares.Logger(middlewares.Autenticar(rota.Funcao))).Methods(rota.Metodo)
	//Ou seja aninhamento de funções

	for _, rota := range rotas {
		if rota.RequerAutenticacao {
			r.HandleFunc(rota.URI, middlewares.Autenticar(rota.Funcao)).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
		}
	}

	return r
}
