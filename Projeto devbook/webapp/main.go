package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookie"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	//Lê as variáveis de ambiente
	config.CarregarVariaveisAmbiente()

	//Cria configura o cookie a ser armazenado no browser.
	cookie.ConfigurarCookie()

	//Criar a requisição que vai chamar a API
	response, erro := http.Post(config.APIAddress("testeapi"), "application/json", nil)
	if erro != nil || response.StatusCode >= 500 {
		if erro == nil {
			erro = errors.New("Servidor da API não está em execução")
		}

		log.Fatal(erro)
	}

	fmt.Println("Web App executando na porta:", config.Porta)
	//Carrega todos os templates
	utils.CarregarTemplates()

	rotas := router.GerarRotasWebAPP()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.Porta), rotas))
}
