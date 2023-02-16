package rotas

import (
	"net/http"
	"webapp/src/router/controllers"
)

var rotasLogin = []WebRotas{
	{
		URI:                "/",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarTelaLogin,
		RequerAutenticacao: false,
	},
	{
		URI:                "/login",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarTelaLogin,
		RequerAutenticacao: false,
	},
}
