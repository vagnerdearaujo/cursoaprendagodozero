package controllers

import (
	"api/banco"
	"api/modelos"
	"api/repositorios"
	"api/src/autenticacao"
	"api/src/resposta"
	"api/src/router/config"
	"api/src/seguranca"
	"api/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}

	//O parâmetro true em validarentidade, garante que a senha será validada
	if erro = usuario.ValidarEntidade(true); erro != nil {
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
	//Obter os parâmetros passados na rota
	parametros := mux.Vars(r) //recebe um map do tipo string
	//A rota foi declarada como: URI: "/usuarios/{usuarioId}"
	id, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		utils.EscreveNaPagina(w, "ObterUsuario: O Id do usuário digitado é inválido")
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}

	var usuario modelos.Usuario
	db, erro := banco.ConectarBanco()
	if erro != nil {
		utils.EscreveNaPagina(w, "Não foi possível conectar ao banco de dados: "+config.StringConexaoBanco)
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()
	repositorioUsuario := repositorios.NovoRepositorioUsuario(db)
	usuario, erro = repositorioUsuario.ObterUsuario(id)
	if erro != nil {
		utils.EscreveNaPagina(w, "Não foi possível recuperar o usuário")
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if usuario.ID == 0 {
		resposta.JSon(w, http.StatusBadRequest, "Usuário não cadastrado")
		return
	}
	resposta.JSon(w, http.StatusOK, usuario)
}

func AlterarUsuario(w http.ResponseWriter, r *http.Request) {
	//Recupera os parâmetros passados na url
	parametros := mux.Vars(r)
	utils.EscreveNaPagina(w, "Alterar dados do Usuário")
	id, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		utils.EscreveNaPagina(w, "O Id do usuário digitado é inválido")
		resposta.JSon(w, http.StatusBadRequest, erro)
		return
	}

	IDUsuarioToken, erro := autenticacao.TokenIDUsuario(r)
	if erro != nil || id != IDUsuarioToken {
		utils.EscreveNaPagina(w, "Usuário não autorizado para esta operação")
		resposta.JSon(w, http.StatusForbidden, erro)
		return
	}

	db, erro := banco.ConectarBanco()
	if erro != nil {
		utils.EscreveNaPagina(w, "Não foi possível conectar ao banco de dados: "+config.StringConexaoBanco)
		resposta.JSon(w, http.StatusUnprocessableEntity, erro)
		return
	}
	defer db.Close()
	repositorioUsuario := repositorios.NovoRepositorioUsuario(db)

	//Verifica se o usuário existe e pode ser recuperado
	usuarioatual, erro := repositorioUsuario.ObterUsuario(id)
	if erro != nil {
		utils.EscreveNaPagina(w, "Não foi possível obter os dados do usuário")
		resposta.JSon(w, http.StatusInternalServerError, erro)
		return
	}

	dadosAlterados, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		utils.EscreveNaPagina(w, "Não foi possível obter os dados do usuário diretamente da página")
		resposta.JSon(w, http.StatusInternalServerError, erro)
		return
	}

	var usuario modelos.Usuario
	if erro := json.Unmarshal(dadosAlterados, &usuario); erro != nil {
		utils.EscreveNaPagina(w, "Não foi possível converter os dados da página em json")
		resposta.JSon(w, http.StatusInternalServerError, erro)
		return
	}

	if erro := usuario.ValidarEntidade(false); erro != nil {
		utils.EscreveNaPagina(w, "Não foi possível transformar dados de usuário em json.")
		resposta.JSon(w, http.StatusBadRequest, erro)
		return
	}
	if (usuario.ID == usuarioatual.ID) && (usuario.Nome == usuarioatual.Nome) && (usuario.Nick == usuarioatual.Nick) && (usuario.Email == usuarioatual.Email) {
		utils.EscreveNaPagina(w, "Nenhum dado foi alterado, atualização não foi realizada.")
		resposta.JSon(w, http.StatusBadRequest, erro)
		return
	}

	if erro := repositorioUsuario.AtualizaDadosUsuario(usuario); erro != nil {
		utils.EscreveNaPagina(w, "Não foi possível atualizar os dados do usuário.")
		resposta.JSon(w, http.StatusInternalServerError, erro)
		return
	}
	resposta.JSon(w, http.StatusOK, "Atualização realizada com sucesso")
}

func ApagarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	id, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		utils.EscreveNaPagina(w, "O Id do usuário digitado é inválido")
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}

	IDUsuarioToken, erro := autenticacao.TokenIDUsuario(r)
	if erro != nil || id != IDUsuarioToken {
		utils.EscreveNaPagina(w, "Usuário não autorizado para esta operação")
		resposta.JSon(w, http.StatusForbidden, erro)
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
	usuarioatual, erro := repositorioUsuario.ObterUsuario(id)
	if erro != nil {
		utils.EscreveNaPagina(w, "Não foi possível obter os dados do usuário")
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	if erro := repositorioUsuario.ApagarUsuario(usuarioatual.ID); erro != nil {
		utils.EscreveNaPagina(w, "Não foi possível excluir os dados do usuário.")
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	resposta.JSon(w, http.StatusOK, "Exclusão realizada com sucesso")
}

func AtualizarSenha(w http.ResponseWriter, r *http.Request) {
	usuarioIDToken, erro := autenticacao.TokenIDUsuario(r)
	if erro != nil {
		resposta.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		resposta.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if usuarioID != usuarioIDToken {
		resposta.Erro(w, http.StatusForbidden, errors.New("Você não pode trocar a senha de outro usuário"))
		return
	}

	requisicaoUsuario, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}

	var senha modelos.Senha
	if erro := json.Unmarshal(requisicaoUsuario, &senha); erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.ConectarBanco()
	if erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorioUsuario := repositorios.NovoRepositorioUsuario(db)
	usuarioBanco, erro := repositorioUsuario.ObterUsuario(usuarioIDToken)
	if erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if usuarioBanco.ID == 0 {
		resposta.Erro(w, http.StatusInternalServerError, errors.New("Este usuário não existe"))
		return
	}

	senhaAtual, erro := repositorioUsuario.ObtemSenhaAtual(usuarioIDToken)
	if erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	/*
		//Este bloco fiz antes de assistir a aula
		//O processo está correto, ainda que eu não tenha usado o pacote de segurança que me permite
		//verificar a senha sem precisar recorrer à validação do usuário

		usuarioBanco.Senha = senha.Nova
		if erro := usuarioBanco.ValidarEntidade(true); erro != nil {
			resposta.Erro(w, http.StatusInternalServerError, erro)
			return
		}

		if senhaAtual != usuarioBanco.Senha {
			resposta.Erro(w, http.StatusInternalServerError, errors.New("Senha atual não confere"))
			return
		}
	*/

	if erro := seguranca.VerificaSenha(senhaAtual, senha.Atual); erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, errors.New("Senha atual não confere"))
		return
	}

	senhaComHash, erro := seguranca.Hash(senha.Nova)
	if erro != nil {
		resposta.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro := repositorioUsuario.AtualizaSenha(usuarioIDToken, string(senhaComHash)); erro != nil {
		resposta.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	resposta.JSon(w, http.StatusOK, "Senha alterada com sucesso")
}
