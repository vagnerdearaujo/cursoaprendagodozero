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
		URI:                "usuarios/{seguidoID}/pararseguir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.PararSeguirUsuario,
		RequerAutenticacao: true,
	},
}