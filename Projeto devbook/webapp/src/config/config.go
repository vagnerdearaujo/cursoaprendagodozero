package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

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

	if HashKey = []byte(os.Getenv("HashKey")); HashKey == nil {

	}

	if BlockKey = []byte(os.Getenv("BlockKey")); BlockKey == nil {

	}

}

func APIAddress(rota string) string {
	return fmt.Sprintf("%v/%v", APIURL, rota)
}
