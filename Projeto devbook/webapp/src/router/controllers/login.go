package controllers

import "net/http"

func CarregarTelaLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Tela de Login.</h1><p><b>Aguarde que vem mais por aí</b></p>"))
}
