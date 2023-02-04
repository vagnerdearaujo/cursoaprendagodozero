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
		RequerAutenticacao: true,
	},
	{
		//Lista todos os usuários
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.ListarUsuarios,
		RequerAutenticacao: true,
	},
	{
		//Obtém o usuário por Id
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.ObterUsuario,
		RequerAutenticacao: true,
	},
	{
		//Altera os dados do usuário
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AlterarUsuario,
		RequerAutenticacao: true,
	},
	{
		//Exclui os dados do usuário
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.ApagarUsuario,
		RequerAutenticacao: false,
	},
}
