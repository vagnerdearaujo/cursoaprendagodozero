package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func CriarUsuarios(w http.ResponseWriter, r *http.Request) {
	//Processa o formulário e deixa as propriedades mais facilmente acessessíveis
	r.ParseForm()

	//Transforma os dados recebidos em um map
	usuarioMap := map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	}

	//cria um json a partir do map de usuários
	usuariojson, erro := json.Marshal(usuarioMap)

	if erro != nil {
		log.Fatal(erro)
	}

	//Criar a requisição que vai chamaar a API
	urlAPI := "http://localhost:5932/usuarios"

	response, erro := http.Post(urlAPI, "application/json", bytes.NewBuffer(usuariojson))
	if erro != nil {
		log.Fatal(erro)
	}
	//O response.Body obrigatoriamente precisa ser fechado.
	//Mesmo que o resultado da API seja nocontent, o body precisa ser fechado.
	defer response.Body.Close()

	fmt.Println(response.Body)
}
