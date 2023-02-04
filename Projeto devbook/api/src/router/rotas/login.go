package rotas

import (
	"api/controllers"
	"net/http"
)

var rotaLogin = Rota{
	URI:                "/login",
	Metodo:             http.MethodPost,
	Funcao:             controllers.LoginUsuario,
	RequerAutenticacao: false,
}
