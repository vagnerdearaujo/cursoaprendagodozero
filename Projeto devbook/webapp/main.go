package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	//Criar a requisição que vai chamaar a API
	urlAPI := "http://localhost:5932/testeapi"

	response, erro := http.Post(urlAPI, "application/json", nil)
	if erro != nil || response.StatusCode >= 500 {
		if erro == nil {
			erro = errors.New("Servidor da API não está em execução")
		}

		log.Fatal(erro)
	}

	porta := "5900"
	fmt.Println("Web App executando na porta:" + porta)
	//Carrega todos os templates
	utils.CarregarTemplates()

	rotas := router.GerarRotasWebAPP()

	log.Fatal(http.ListenAndServe(":"+porta, rotas))
}
