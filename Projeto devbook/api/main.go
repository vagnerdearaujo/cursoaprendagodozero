package main

import (
	"api/src/router"
	"api/src/router/config"
	"fmt"
	"log"
	"net/http"
)

func main() {

	/*
		//Quando precisar gerar um novo secret key aleatório, descomente este bloco
		fmt.Println(utils.GeraSecretKey())
		return
	*/
	config.InicializaAmbiente()
	fmt.Println(string(config.API_SecretKey))
	fmt.Printf("Escutando API na porta %s\nBanco de Dados: %s\nServidor:%s\n", config.PortaAPI, config.DB_nome, config.DB_server)

	r := router.GerarRotas()

	log.Fatal(http.ListenAndServe(":"+config.PortaAPI, r))
}
