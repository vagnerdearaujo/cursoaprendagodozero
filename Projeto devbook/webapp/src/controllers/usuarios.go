package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/respostas"
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
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//Criar a requisição que vai chamaar a API
	urlAPI := config.APIAddress("usuarios")

	response, erro := http.Post(urlAPI, "application/json", bytes.NewBuffer(usuariojson))
	if erro != nil {
		respostas.JSON(w, http.StatusBadGateway, respostas.ErroAPI{Erro: fmt.Sprintf("O Servidor %v não respondeu a requisição", urlAPI)})
		return
	}

	//O response.Body obrigatoriamente precisa ser fechado.
	//Mesmo que o resultado da API seja nocontent, o body precisa ser fechado.
	defer response.Body.Close()

	/*
		Neste ponto se a requisição foi executada, o erro será nil, uma vez que o erro refere-se exclusivamente
		ao envio e recebimento de resposta da requisilção e não ao status code.
		Se o statuscode retornar da API como 500 (internal error) do ponto de vista da chamada à API terá havido sucesso.
		Para capturar se a API conseguiu ou não processar adequadamente a requisição, o response.StatusCode será
		utilizado para gerar o erro a ser enviado para a página.

		https://www.httpstatus.com.br/
		consultado em: 19/02/2023

		1xx Informativo
		200 OK
		3xx Redirecionamento
		4xx Erro no Cliente
		5xx Erro no Servidor

		Os códigos individiuais de cada categoria podem ser obtidos no arquivo   ../../ideias/StatusCodeHTTP.txt
		Ou no endereço: https://www.httpstatus.com.br/
	*/

	//Códigos de falha estão na "classe 400" e na "classe 500"
	if response.StatusCode >= 400 {
		respostas.TratarStatusCode(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	respostas.JSON(w, http.StatusOK, "Usuário Logado com sucesso")
}