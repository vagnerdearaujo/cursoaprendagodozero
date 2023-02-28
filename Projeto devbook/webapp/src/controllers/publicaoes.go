package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/requisicoes"
	"webapp/src/respostas"

	"github.com/gorilla/mux"
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

func CurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoId, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	urlAPI := config.APIAddress(fmt.Sprintf("publicacoes/%d/curtir", publicacaoId))

	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, urlAPI, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusBadGateway, respostas.ErroAPI{Erro: fmt.Sprintf("O Servidor %v não respondeu a requisição\n%s", urlAPI, erro.Error())})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCode(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}

func DescurtirPublicacao(w http.ResponseWriter, r *http.Request) {
}

func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {
}

func ExcluirPublicacao(w http.ResponseWriter, r *http.Request) {
}
