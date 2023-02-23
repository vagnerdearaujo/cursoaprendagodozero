package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gorilla/securecookie"
	"github.com/joho/godotenv"
)

var (
	//URL da API
	APIURL = ""

	//Porta da API
	Porta = 0

	//HaskKey é utilizado para autenticar o cookie
	HashKey []byte

	//BlockKey é utilizado para criptografar o cookie
	BlockKey []byte

	//Nome do Cookie da aplicação
	CookieName = ""
)

func CarregarVariaveisAmbiente() {
	var erro error
	if erro := godotenv.Load(); erro != nil {
		log.Fatal(errors.New("Arquivo de parâmetros de configuração não encontrado."))
	}

	if Porta, erro = strconv.Atoi(os.Getenv("Porta")); erro != nil {
		Porta = 5900
	}

	if APIURL = os.Getenv("APIURL"); APIURL == "" {
		log.Fatal(errors.New("URL da API não definida"))
	}

	if CookieName = os.Getenv("CookieName"); CookieName == "" {
		CookieName = "devbook"
	}

	//Na versão original o arquivo ambiente possui as chaves, nesta versão estou substituindo
	//por geração randômica para cada vez que o projeto for executado

	HashKey = securecookie.GenerateRandomKey(32)
	BlockKey = securecookie.GenerateRandomKey(32)
}

func APIAddress(rota string) string {
	return fmt.Sprintf("%v/%v", APIURL, rota)
}
