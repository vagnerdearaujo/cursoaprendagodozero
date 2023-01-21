package controllers

import (
	"api/utils"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	utils.EscreveNaPagina(w, "Criando Usuário")
}

func ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	utils.EscreveNaPagina(w, "Listar todos os Usuários")
}

func ObterUsuario(w http.ResponseWriter, r *http.Request) {
	utils.EscreveNaPagina(w, "Obter um Usuário por ID")
}

func AlterarUsuario(w http.ResponseWriter, r *http.Request) {
	utils.EscreveNaPagina(w, "Alterar dados do Usuário")
}

func ApagarUsuario(w http.ResponseWriter, r *http.Request) {
	utils.EscreveNaPagina(w, "Apagar dados do Usuário")
}
