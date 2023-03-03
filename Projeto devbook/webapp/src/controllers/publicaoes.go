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

// NovaPublicacao: Chama a api para inclusão de nova publicação no banco de dados
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

// CurtirPublicacao: Chama a api para informar uma curtida
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

// DescurtirPublicacao: Chama a api para informar uma descurtida
func DescurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoId, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	urlAPI := config.APIAddress(fmt.Sprintf("publicacoes/%d/descurtir", publicacaoId))

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

// AtualizarPublicacao: Captura os dados da requisição e os envia para a API para atualização no banco
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoId, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

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
	urlAPI := config.APIAddress(fmt.Sprintf("publicacoes/%d", publicacaoId))

	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPut, urlAPI, bytes.NewBuffer(publicacaojson))
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

// ExcluirPublicacao: Chama a api para excluir uma publicação do banco de dados
func ExcluirPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoId, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//Criar a requisição que vai chamar a API
	urlAPI := config.APIAddress(fmt.Sprintf("publicacoes/%d", publicacaoId))

	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodDelete, urlAPI, nil)
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
