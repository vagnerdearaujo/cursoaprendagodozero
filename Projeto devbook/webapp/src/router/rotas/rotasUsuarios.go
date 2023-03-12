package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasUsuarios = []WebRotas{
	{
		URI:                "/criar-usuario",
		Funcao:             controllers.CarregarPaginaCadastroUsuario,
		Metodo:             http.MethodGet,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios",
		Funcao:             controllers.CriarUsuarios,
		Metodo:             http.MethodPost,
		RequerAutenticacao: false,
	},
	{
		URI:                "/home",
		Funcao:             controllers.Home,
		Metodo:             http.MethodGet,
		RequerAutenticacao: true,
	},
	{
		URI:                "/buscar-usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaUsuarios,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPerfilUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuarioId}/pararseguir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.PararSeguir,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuarioId}/seguir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.Seguir,
		RequerAutenticacao: true,
	},
}
