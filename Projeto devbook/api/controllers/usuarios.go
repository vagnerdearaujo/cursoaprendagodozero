package controllers

import (
	"api/banco"
	"api/modelos"
	"api/repositorios"
	"api/src/resposta"
	"api/src/router/config"
	"api/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		utils.EscreveNaPagina(w, "Não foi possível carregar dados da página.")
		resposta.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario

	if erro := json.Unmarshal(corpoRequest, &usuario); erro != nil {
		utils.EscreveNaPagina(w, "Não foi possível transformar dados de usuário em json.")
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar(); erro != nil {
		utils.EscreveNaPagina(w, "Não foi possível transformar dados de usuário em json.")
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.ConectarBanco()
	if erro != nil {
		utils.EscreveNaPagina(w, "Não foi possível conectar ao banco de dados: "+config.StringConexaoBanco)
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorioUsuario := repositorios.NovoRepositorioUsuario(db)
	usuario.ID, erro = repositorioUsuario.NovoUsuario(usuario)
	if erro != nil {
		utils.EscreveNaPagina(w, "Erro ao tentar incluir novo usuário")
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	utils.EscreveNaPagina(w, fmt.Sprintf("Usuário incluído com sucesso: %d", usuario.ID))

	//Devolve o JSon do usuário atualizado com o ID e sem a data de criação
	//Corrigir para chamar o método de consulta por ID para retornar inclusive com a data e hora de criação
	resposta.JSon(w, http.StatusCreated, usuario)

}

func ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	var usuarios []modelos.Usuario

	db, erro := banco.ConectarBanco()
	if erro != nil {
		utils.EscreveNaPagina(w, "Não foi possível conectar ao banco de dados: "+config.StringConexaoBanco)
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repositorioUsuarios := repositorios.NovoRepositorioUsuario(db)
	//Captura o parâmetro usuario passado na url
	nickouname := strings.ToLower(r.URL.Query().Get("usuario"))
	usuarios, erro = repositorioUsuarios.ListarUsuarios(nickouname)
	if erro != nil {
		utils.EscreveNaPagina(w, "Não foi possível conectar ao banco de dados: "+config.StringConexaoBanco)
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	resposta.JSon(w, http.StatusOK, usuarios)
}

func ObterUsuario(w http.ResponseWriter, r *http.Request) {
	utils.EscreveNaPagina(w, "Obter um Usuário por ID")
}

func AlterarUsuario(w http.ResponseWriter, r *http.Request) {
	utils.EscreveNaPagina(w, "Alterar dados do Usuário")
}

func ApagarUsuario(w http.ResponseWriter, r *http.Request) {
	utils.EscreveNaPagina(w, "Apagar dados do Usuário")
}
