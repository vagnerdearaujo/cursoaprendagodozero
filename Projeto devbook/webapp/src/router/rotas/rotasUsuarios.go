package rotas

import (
	"net/http"
	"webapp/src/router/controllers"
)

var rotasUsuarios = []WebRotas{
	{
		URI:                "/criar-usuario",
		Funcao:             controllers.CarregarPaginaCadastroUsuario,
		Metodo:             http.MethodGet,
		RequerAutenticacao: false,
	},
}
