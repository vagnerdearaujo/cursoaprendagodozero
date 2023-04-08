package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookie"
	"webapp/src/modelos"
	"webapp/src/respostas"
	"webapp/src/utils"
)

// AutenticarUsuario autentica o usuário usando e-mail e senha
func AutenticarUsuario(w http.ResponseWriter, r *http.Request) {

	//Processa o formulário
	r.ParseForm()

	//Transforma os dados da requisição em map
	login := map[string]string{
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	}

	loginjson, erro := json.Marshal(login)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, erro)
		return
	}

	//Criar a requisição que vai chamaar a API
	urlAPI := config.APIAddress("login")

	response, erro := http.Post(urlAPI, "application/json", bytes.NewBuffer(loginjson))
	if erro != nil {
		//Neste ponto não se pode usar o response.StatusCode, porque se houve erro o response será nil.
		//Se o response for igual a nil será levantada uma exceção fatal Panic.
		respostas.JSON(w, http.StatusBadGateway, respostas.ErroAPI{Erro: fmt.Sprintf("O Servidor %v não respondeu a requisição", urlAPI)})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCode(w, response)
		return
	}

	var dadosAutenticacao modelos.DadosAutenticacao
	if erro := json.NewDecoder(response.Body).Decode(&dadosAutenticacao); erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	if erro := cookie.ArmazenaCookie(w, dadosAutenticacao.ID, dadosAutenticacao.Token); erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}

func Perfil(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "underconstruction.html", nil)
}

func FazerLogout(w http.ResponseWriter, r *http.Request) {
	cookie.LimparCookie(w)
	//Chamar com o redirect evita que no topo da página fique aparecendo a rota logout e ainda informa do redirecionamento via status 302
	http.Redirect(w, r, "/login", http.StatusPermanentRedirect)

	//O utils.ExecutarTemplate resolve o problema, no entanto na barra de endereço fica a rota /logout
	//utils.ExecutarTemplate(w, "login.html", nil)
}
