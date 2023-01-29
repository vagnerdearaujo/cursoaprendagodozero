package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConexaoBanco = ""
	PortaAPI           = ""
	DB_usuario         = ""
	DB_senha           = ""
	DB_nome            = ""
	DB_server          = ""
	DB_driverbanco     = ""
)

// InicializaAmbiente carrega as variáveis de ambiente

func InicializaAmbiente() {
	var (
		erro          error
		db_parametros string
	)

	//A função Load carrega valores para as variáveis de ambiente conforme definição
	//do arquivo .env
	if erro = godotenv.Load(); erro != nil {
		log.Fatal("Arquivo de configuração de ambiente não encontrado." + erro.Error())
	}

	//Todas os valores declarados no .env são strings, portanto, no caso da porta
	//é necessário a conversão para int.

	//A função Getenv lê uma variável de ambiente do OS
	porta, erro := strconv.Atoi(os.Getenv("API_PORTA"))
	if erro != nil {
		porta = 5932
	}
	PortaAPI = strconv.Itoa(porta)

	DB_usuario = os.Getenv("DB_USUARIO")
	if DB_usuario == "" {
		fmt.Println("*** Erro: Nome do usuário do banco não definido. ****")
		log.Fatal("Nome do usuário do banco não definido." + erro.Error())
	}

	DB_senha = os.Getenv("DB_SENHA")
	if DB_senha == "" {
		fmt.Println("*** Senha do usuário do banco não definida. ****")
		log.Fatal("Senha do usuário do banco não definida." + erro.Error())
	}

	DB_nome = os.Getenv("DB_NOME")
	if DB_nome == "" {
		fmt.Println("*** Nome do banco não definido. ****")
		log.Fatal("Nome do banco não definido." + erro.Error())
	}

	DB_server = os.Getenv("DB_SERVIDOR")
	if DB_server == "" {
		fmt.Println("*** Servidor do banco não definido. ****")
		log.Fatal("Servidor do banco não definido." + erro.Error())
	}

	DB_driverbanco = os.Getenv("DB_DRIVERBANCO")
	if DB_driverbanco == "" {
		DB_driverbanco = "mysql"
	}

	db_parametros = os.Getenv("DB_PARAMETROS")
	if db_parametros == "" {
		db_parametros = "?charset=utf8&parseTime=True&loc=Local"
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@tcp(%s)/%s%s", DB_usuario, DB_senha, DB_server, DB_nome, db_parametros)
}
