package controllers

import (
	"api/banco"
	"api/modelos"
	"api/repositorios"
	"api/src/autenticacao"
	"api/src/resposta"
	"encoding/json"
	"errors"
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

	if erro := publicao.Preparar(); erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
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
	usuarioId, erro := autenticacao.TokenIDUsuario(r)
	if erro != nil {
		resposta.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	publicaoId, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		resposta.Erro(w, http.StatusForbidden, erro)
		return
	}

	db, erro := banco.ConectarBanco()
	if erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorioPublicacao := repositorios.NovoRepositorioPublicacao(db)
	publicacaoBanco, erro := repositorioPublicacao.ObterPublicacaoPorId(publicaoId)
	if erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if publicacaoBanco.AutorID != usuarioId {
		resposta.Erro(w, http.StatusForbidden, errors.New("Você não pode alterar publicação de outro usuário."))
		return
	}

	corpoPublicacao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		resposta.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacao modelos.Publicacao
	if erro := json.Unmarshal(corpoPublicacao, &publicacao); erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro := publicacao.Preparar(); erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}

	publicacao.ID = publicaoId

	if erro := repositorioPublicacao.AtualizarPublicacao(publicacao); erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	resposta.JSon(w, http.StatusNoContent, nil)

}

func ExcluirPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioId, erro := autenticacao.TokenIDUsuario(r)
	if erro != nil {
		resposta.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	publicaoId, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		resposta.Erro(w, http.StatusForbidden, erro)
		return
	}

	db, erro := banco.ConectarBanco()
	if erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorioPublicacao := repositorios.NovoRepositorioPublicacao(db)
	publicacaoBanco, erro := repositorioPublicacao.ObterPublicacaoPorId(publicaoId)
	if erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if publicacaoBanco.AutorID != usuarioId {
		resposta.Erro(w, http.StatusForbidden, errors.New("Você não pode excluir publicação de outro usuário."))
		return
	}

	if erro := repositorioPublicacao.ExcluirPublicacao(publicaoId); erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	resposta.JSon(w, http.StatusNoContent, nil)
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
	publicacao, erro := repositoriosPublicacao.ObterPublicacaoPorId(publicacoId)
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

func BuscarPublicacaoUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioId, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.ConectarBanco()
	if erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repositorioPublicacao := repositorios.NovoRepositorioPublicacao(db)
	if erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	publicacoes, erro := repositorioPublicacao.ObterPublicacaoPorUsuario(usuarioId)
	if erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	resposta.JSon(w, http.StatusOK, publicacoes)
}

func CurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	statuscode, erro := curtirdescurtirpublicacao(r, true)
	if erro != nil {
		resposta.Erro(w, statuscode, erro)
		return
	}

	resposta.JSon(w, http.StatusNoContent, nil)
}

func DescurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	statuscode, erro := curtirdescurtirpublicacao(r, false)
	if erro != nil {
		resposta.Erro(w, statuscode, erro)
		return
	}

	resposta.JSon(w, http.StatusNoContent, nil)

}

func curtirdescurtirpublicacao(r *http.Request, curtir bool) (int, error) {
	parametros := mux.Vars(r)
	publicacoId, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		return http.StatusBadRequest, erro
	}

	db, erro := banco.ConectarBanco()
	if erro != nil {
		return http.StatusInternalServerError, erro
	}
	defer db.Close()

	repositoriosPublicacao := repositorios.NovoRepositorioPublicacao(db)
	publicacao, erro := repositoriosPublicacao.ObterPublicacaoPorId(publicacoId)
	if erro != nil {
		return http.StatusBadRequest, erro
	}

	if curtir {
		publicacao.Curtidas = publicacao.Curtidas + 1
	} else {
		publicacao.Curtidas = publicacao.Curtidas - 1
	}

	if erro := publicacao.Preparar(); erro != nil {
		return http.StatusBadRequest, erro
	}

	if erro := repositoriosPublicacao.CurtirDescuPublicacao(publicacao); erro != nil {
		return http.StatusInternalServerError, erro
	}
	return http.StatusOK, nil
}
