package controllers

import (
	"api/banco"
	"api/modelos"
	"api/repositorios"
	"api/src/autenticacao"
	"api/src/resposta"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	idUsuario, erro := autenticacao.TokenIDUsuario(r)
	if erro != nil {
		resposta.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	corpoPublicacao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}

	var publicao modelos.Publicacao
	if erro := json.Unmarshal(corpoPublicacao, &publicao); erro != nil {
		resposta.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := banco.ConectarBanco()
	if erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorioPublicacao := repositorios.NovoRepositorioPublicacao(db)

	publicao.AutorID = idUsuario

	publicacaoId, erro := repositorioPublicacao.IncluirrPublicacao(publicao)
	if erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	resposta.JSon(w, http.StatusCreated, fmt.Sprintf("Publicação realizada com sucesso: %v", publicacaoId))
}
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {

}
func ExcluirPublicacao(w http.ResponseWriter, r *http.Request) {

}
func BuscarPublicacaoId(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacoId, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.ConectarBanco()
	if erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositoriosPublicacao := repositorios.NovoRepositorioPublicacao(db)
	publicacao, erro := repositoriosPublicacao.ListarPublicacaoId(publicacoId)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}

	resposta.JSon(w, http.StatusOK, publicacao)

}
func ListarPublicacoes(w http.ResponseWriter, r *http.Request) {
	autorId, erro := autenticacao.TokenIDUsuario(r)
	if erro != nil {
		resposta.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	db, erro := banco.ConectarBanco()
	if erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorioPublicacao := repositorios.NovoRepositorioPublicacao(db)
	publicacoes, erro := repositorioPublicacao.ListarPublicacoes(autorId)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}
	resposta.JSon(w, http.StatusOK, publicacoes)
}
