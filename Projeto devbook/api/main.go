package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

const ApiPorta = "5932"

func main() {
	fmt.Println("Escutando API na porta " + ApiPorta)

	r := router.GerarRotas()

	log.Fatal(http.ListenAndServe(":"+ApiPorta, r))
}
