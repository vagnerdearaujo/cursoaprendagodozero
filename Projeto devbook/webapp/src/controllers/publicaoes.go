package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
)

func NovaPublicacao(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	publicacaoMap := map[string]string{
		"titulo":   r.FormValue("titulo"),
		"conteudo": r.FormValue("conteudo"),
	}

	publicacaojson, erro := json.Marshal(publicacaoMap)

	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//Criar a requisição que vai chamar a API
	urlAPI := config.APIAddress("publicacoes")

	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, urlAPI, bytes.NewBuffer(publicacaojson))
	if erro != nil {
		respostas.JSON(w, http.StatusBadGateway, respostas.ErroAPI{Erro: fmt.Sprintf("O Servidor %v não respondeu a requisição", urlAPI)})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCode(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}
