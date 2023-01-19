package main

import (
	"crud/banco"
	"crud/servidor"
	"crud/settings"
	"fmt"
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
	mysqlSettings := settings.MySQLSettings()

	db, erro := banco.ConectarDB(mysqlSettings)

	if erro != nil {
		fmt.Println(erro)
		return
	}
	fmt.Printf("Conexão realizada com sucesso (Servidor:%s / Banco:%s)\n", mysqlSettings.DBDriver, mysqlSettings.DBName)

	defer db.Close()

	if erro != nil {
		log.Fatal(erro)
	}

	//Ainda que a sintaxe seja igual ao http.HandleFunc o método mux permite
	//especificar o verbo http
	router.HandleFunc("/usuario", servidor.CriarUsuario).Methods(http.MethodPost)

	//Ainda que a rota seja a mesma, não haverá problema algum quando os verbos forem direntes
	router.HandleFunc("/usuario", servidor.ListarUsuarios).Methods(http.MethodGet)

	//Quando a rota possui um parâmetro variável, por exemplo Id, este deve ser colocado na rota entre {}
	router.HandleFunc("/usuario/{id}", servidor.BuscarUsuario).Methods(http.MethodGet)
	router.HandleFunc("/usuario/{id}", servidor.AtualizarUsuario).Methods(http.MethodPut)
	router.HandleFunc("/usuario/{id}", servidor.ExcluirUsuario).Methods(http.MethodDelete)

	println("Servidor ativo e escutando a porta: " + mysqlSettings.ServerPort)

	log.Fatal(http.ListenAndServe(":"+mysqlSettings.ServerPort, router))
}
