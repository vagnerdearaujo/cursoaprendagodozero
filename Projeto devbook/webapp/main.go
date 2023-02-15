package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
)

func main() {
	porta := "5900"
	fmt.Println("Web App executando na porta:" + porta)
	rotas := router.GerarRotasWebAPP()
	
	log.Fatal(http.ListenAndServe(":"+porta, rotas))
}
