package main

import (
	"crud/banco"
	"crud/servidor"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*
	Para o crud será utilizado um pacote mais robusto que o http.
	go get github.com/gorilla/mux

	O mux permite ter um gerenciamento mais robusto para as rotas, ainda
	que o pacote http continue sendo utilizado para levantar o servidor.

	A criação de rotas é feita com o mux.NewRouter
*/

/*
CRUD (Create, Read, Update & Delete)
Create	= Post
Read	= Get
Update	= Put
Delete	= Delete
*/
func main() {
	router := mux.NewRouter()

	//Sobe o servidor http, informando o gerenciador de rotas.
	serverPort := "5932"
	dbUser := "golang_devbook"
	dbPasswd := "devbook_golang"
	dbDriver := "mysql"
	dbName := "devbook"
	dbServer := "172.18.0.2:3306"
	connectionParameters := "?charset=utf8&parseTime=True&loc=Local"
	db, erro := banco.ConectarDB(dbDriver, dbUser, dbPasswd, dbName, dbServer, connectionParameters)
	defer db.Close()
	if erro != nil {
		log.Fatal(erro)
	}

	//Ainda que a sintaxe seja igual ao http.HandleFunc o método mux permite
	//especificar o verbo http
	router.HandleFunc("/usuario", servidor.CriarUsuario).Methods(http.MethodPost)
	println("Servidor ativo e escutando a porta: " + serverPort)

	log.Fatal(http.ListenAndServe(":"+serverPort, router))
}
