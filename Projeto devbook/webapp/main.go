package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	porta := "5900"
	fmt.Println("Web App executando na porta:" + porta)
	//Carrega todos os templates
	utils.CarregarTemplates()

	rotas := router.GerarRotasWebAPP()

	log.Fatal(http.ListenAndServe(":"+porta, rotas))
}
