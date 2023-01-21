package rotas

import (
	"api/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		//Inclusão de usuários
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		//Lista todos os usuários
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.ListarUsuarios,
		RequerAutenticacao: false,
	},
	{
		//Obtém o usuário por Id
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.ObterUsuario,
		RequerAutenticacao: false,
	},
	{
		//Altera os dados do usuário
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AlterarUsuario,
		RequerAutenticacao: false,
	},
	{
		//Exclui os dados do usuário
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.ApagarUsuario,
		RequerAutenticacao: false,
	},
}
