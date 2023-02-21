package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	config.CarregarVariaveisAmbiente()
	//Criar a requisição que vai chamaar a API

	response, erro := http.Post(config.APIAddress("/testeapi"), "application/json", nil)
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
