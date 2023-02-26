package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookie"
	"webapp/src/modelos"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
	"webapp/src/utils"
)

func CarregarTelaLogin(w http.ResponseWriter, r *http.Request) {
	//Carrega a página do login
	utils.ExecutarTemplate(w, "login.html", nil)
}

func CarregarPaginaCadastroUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastrousuarios.html", nil)
}

func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	/*
		Diferentemente das páginas de login e cadastro de usuários,
		a página home necessita de ser chamada com o uso de um token, portanto a
		chamada: utils.ExecutarTemplate(w, "home.html", nil) irá ser recusada pela API,
		justamente por não possuir um token.
	*/

	/*
		O código abaixo se executado, irá imprimir no terminal o status 401 (Não autorizado), justamente
		porque a API não recebeu o token, portanto, devolvendo o status 401.

		response, erro := http.Get(config.APIAddress("publicacoes"))
		fmt.Println(response.StatusCode, erro)
		utils.ExecutarTemplate(w, "home.html", nil)
	*/

	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, config.APIAddress("publicacoes"), nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCode(w, response)
		return
	}

	//Captura todas as publicações recebidas da API
	var publicacoes []modelos.Publicacao

	//Popula publicacoes com os dados recebidos no response.
	if erro := json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//Posso ignorar o erro, porque se chegou aqui o cookie existe
	cookie, _ := cookie.CarregarCookie(r)

	//Sabendo que a API sempre vai passar um número para o cookie
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	//Poderia ter criado a struct diretamente na linha de chamada de ExecutarTemplate,
	//porém acho que o código fica muito confuso
	type dadosPublicacao struct {
		UsuarioID   uint64
		Publicacoes []modelos.Publicacao
	}

	utils.ExecutarTemplate(w, "home.html", dadosPublicacao{
		UsuarioID:   usuarioID,
		Publicacoes: publicacoes,
	})
}
