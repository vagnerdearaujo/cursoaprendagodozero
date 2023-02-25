package rotas

import (
	"net/http"
	"webapp/src/controllers"
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
	{
		URI:                "/login",
		Metodo:             http.MethodPost,
		Funcao:             controllers.AutenticarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/perfil",
		Metodo:             http.MethodGet,
		Funcao:             controllers.Perfil,
		RequerAutenticacao: false,
	},
	{
		URI:                "/logout",
		Metodo:             http.MethodGet,
		Funcao:             controllers.Logout,
		RequerAutenticacao: false,
	},
}
