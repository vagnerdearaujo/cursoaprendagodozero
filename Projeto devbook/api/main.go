package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Escutando API na porta 5932")

	r := router.GerarRotas()

	log.Fatal(http.ListenAndServe(":5932", r))
}
