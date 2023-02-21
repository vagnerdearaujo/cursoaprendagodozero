package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"webapp/src/respostas"
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
	urlAPI := "http://localhost:5932/login"

	response, erro := http.Post(urlAPI, "application/json", bytes.NewBuffer(loginjson))
	if erro != nil {
		respostas.JSON(w, http.StatusBadGateway, respostas.ErroAPI{Erro: fmt.Sprintf("O Servidor %v não respondeu a requisição", urlAPI)})
		return
	}

	defer response.Body.Close()

	token, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(token))

	if response.StatusCode >= 400 {
		respostas.TratarStatusCode(w, response)
		return
	}
	respostas.JSON(w, response.StatusCode, nil)

}
