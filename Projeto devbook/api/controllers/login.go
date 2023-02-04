package controllers

import (
	"api/banco"
	"api/modelos"
	"api/repositorios"
	"api/src/autenticacao"
	"api/src/resposta"
	"api/src/seguranca"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func LoginUsuario(w http.ResponseWriter, r *http.Request) {
	var usuarioLogin modelos.Usuario
	usuarioBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		resposta.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	erro = json.Unmarshal(usuarioBody, &usuarioLogin)
	if erro != nil {
		resposta.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	db, erro := banco.ConectarBanco()
	if erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	usuarioRepositorio := repositorios.NovoRepositorioUsuario(db)

	credencialUsuario, erro := usuarioRepositorio.ObterUsuarioPorEmail(usuarioLogin.Email)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if credencialUsuario.ID == 0 {
		resposta.Erro(w, http.StatusUnauthorized, errors.New("Usuário não cadastrado ou senha incorreta"))
		return
	}

	usuarioBanco, erro := usuarioRepositorio.ObterUsuario(credencialUsuario.ID)
	if erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro := seguranca.VerificaSenha(credencialUsuario.Senha, usuarioLogin.Senha); erro != nil {
		resposta.Erro(w, http.StatusUnauthorized, errors.New("Usuário não cadastrado ou senha incorreta"))
		return
	}

	resposta.JSon(w, http.StatusAccepted, "Login realizado com sucesso: "+usuarioBanco.Nome)

	token, _ := autenticacao.CriarToken(credencialUsuario.ID)
	resposta.JSon(w, http.StatusOK, token)
}
