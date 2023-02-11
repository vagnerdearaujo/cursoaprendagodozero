package controllers

import (
	"api/banco"
	"api/modelos"
	"api/repositorios"
	"api/src/autenticacao"
	"api/src/resposta"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.TokenIDUsuario(r)
	if erro != nil {
		resposta.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	seguidoID, erro := strconv.ParseUint(parametros["seguidoID"], 10, 64)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if seguidoID == usuarioID {
		resposta.Erro(w, http.StatusForbidden, errors.New("Não é permitido seguir a si mesmo"))
		return
	}

	db, erro := banco.ConectarBanco()
	if erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repositorioUsuario := repositorios.NovoRepositorioUsuario(db)
	usuarioSeguido, erro := repositorioUsuario.SeguirUsuario(usuarioID, seguidoID)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}
	resposta.JSon(w, http.StatusOK, "Agora você segue: "+usuarioSeguido.Nome)

}

func PararSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	seguidorID, erro := autenticacao.TokenIDUsuario(r)
	if erro != nil {
		resposta.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parameters := mux.Vars(r)
	seguidoID, erro := strconv.ParseUint(parameters["seguidoID"], 10, 64)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if seguidoID == seguidorID {
		resposta.Erro(w, http.StatusForbidden, errors.New("Você não pode deixar de seguir você mesmo."))
		return
	}

	db, erro := banco.ConectarBanco()
	if erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	usuarioRepositorio := repositorios.NovoRepositorioUsuario(db)
	usuarioSeguido, erro := usuarioRepositorio.ObterUsuario(seguidoID)
	if erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	seguia, erro := usuarioRepositorio.PararSeguir(seguidorID, seguidoID)
	if erro != nil {
		resposta.Erro(w, http.StatusForbidden, erro)
		return
	}

	if !seguia {
		resposta.Erro(w, http.StatusBadRequest, errors.New("Você não seguia o usuário: "+usuarioSeguido.Nome))
		return
	}

	resposta.JSon(w, http.StatusOK, "Você não segue mais: "+usuarioSeguido.Nome)
}

func MeusSeguidores(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.TokenIDUsuario(r)
	if erro != nil {
		resposta.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}
	seguidores, erro := ObtemSeguidores(usuarioID)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}
	resposta.JSon(w, http.StatusOK, seguidores)
}

func SeguidosPorMim(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.TokenIDUsuario(r)
	if erro != nil {
		resposta.Erro(w, http.StatusUnauthorized, erro)
	}

	seguidos, erro := ObtemSeguidos(usuarioID)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
	}

	resposta.JSon(w, http.StatusOK, seguidos)
}

func Seguidores(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioID"], 10, 64)

	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}
	seguidores, erro := ObtemSeguidores(usuarioID)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}
	resposta.JSon(w, http.StatusOK, seguidores)
}

func ObtemSeguidores(usuarioId uint64) ([]modelos.Usuario, error) {
	db, erro := banco.ConectarBanco()
	if erro != nil {
		return nil, erro
	}
	defer db.Close()
	repositorioUsuario := repositorios.NovoRepositorioUsuario(db)
	seguidores, erro := repositorioUsuario.ObterSeguidores(usuarioId)
	if erro != nil {
		return nil, erro
	}

	if seguidores == nil {
		return nil, errors.New("Usuário informado não tem seguidores")
	}

	return seguidores, nil
}

func SeguidoPor(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioID"], 10, 64)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}
	seguidos, erro := ObtemSeguidos(usuarioID)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}

	resposta.JSon(w, http.StatusOK, seguidos)
}

func ObtemSeguidos(usuarioID uint64) ([]modelos.Usuario, error) {
	db, erro := banco.ConectarBanco()
	if erro != nil {
		return nil, erro
	}
	defer db.Close()

	repositorioUsuario := repositorios.NovoRepositorioUsuario(db)
	seguidos, erro := repositorioUsuario.ObterSeguidos(usuarioID)
	if erro != nil {
		return nil, erro
	}
	if seguidos == nil {
		return nil, errors.New("Ninguém é seguido pelo usuário informado")
	}

	return seguidos, nil
}
