package rotas

import (
	"api/controllers"
	"net/http"
)

var RotaSeguidores = []Rota{
	{
		URI:                "/usuarios/{seguidoID}/seguir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.SeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{seguidoID}/pararseguir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.PararSeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/meusseguidores",
		Metodo:             http.MethodGet,
		Funcao:             controllers.MeusSeguidores,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/seguidospormim",
		Metodo:             http.MethodGet,
		Funcao:             controllers.SeguidosPorMim,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuarioID}/seguidores",
		Metodo:             http.MethodGet,
		Funcao:             controllers.Seguidores,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuarioID}/seguidopor",
		Metodo:             http.MethodGet,
		Funcao:             controllers.SeguidoPor,
		RequerAutenticacao: true,
	},
}
