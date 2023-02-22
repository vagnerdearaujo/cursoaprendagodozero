package controllers

import (
	"net/http"
	"webapp/src/utils"
)

func CarregarTelaLogin(w http.ResponseWriter, r *http.Request) {
	//Carrega a página do login
	utils.ExecutarTemplate(w, "login.html", nil)
}

func CarregarPaginaCadastroUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastrousuarios.html", nil)
}

func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "home.html", nil)
}
