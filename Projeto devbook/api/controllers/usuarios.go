package controllers

import (
	"api/banco"
	"api/modelos"
	"api/repositorios"
	"api/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		utils.EscreveNaPagina(w, "Não foi possível carregar dados da página.")
		return
	}

	var usuario modelos.Usuario

	if erro := json.Unmarshal(corpoRequest, &usuario); erro != nil {
		utils.EscreveNaPagina(w, "Não foi possível transformar dados de usuário em json.")
		return
	}

	db, erro := banco.ConectarBanco()
	if erro != nil {
		utils.EscreveNaPagina(w, "Não foi possível conectar ao banco de dados")
		return
	}
	defer db.Close()

	repositorioUsuario := repositorios.NovoRepositorioUsuario(db)
	id, erro := repositorioUsuario.NovoUsuario(usuario)
	if erro != nil {
		utils.EscreveNaPagina(w, "Erro ao tentar incluir novo usuário")
	}
	utils.EscreveNaPagina(w, fmt.Sprintf("Usuário incluído com sucesso: %d", id))

}

func ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	utils.EscreveNaPagina(w, "Listar todos os Usuários")
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
