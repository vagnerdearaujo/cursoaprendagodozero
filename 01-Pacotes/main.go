package main

/*
	Todo executável obrigatoriamente deve ser declarado no package main e possuir uma função main.
	O comando go mod init <nome_módulo> cria um arquivo go.mod com as configurações de dependências.
	Um pacote é composto por 1 ou mais fontes em go e que estejam em único diretório e compartilhem a mesma diretiva package
	Tudo o que for declarado em um fonte de um package é visível em todos os fontes deste mesmo package.
	Uma variável ou função é considerada "pública" se o nome for iniciado por uma letra maíuscula

	Para incluir um pacote externo: go get <url do pacote>
	exemplo: go get github.com/badoux/checkmail

	Uma linha será incluída no arquivo go.mod, indicando esta dependência
	require github.com/badoux/checkmail v1.2.1 // indirect

*/

import (
	"fmt"
	"modulo/auxiliar"
)

func main() {
	fmt.Println("=========================== Pacotes ===========================")

	auxiliar.Escrever()
	fmt.Println("Escrevendo do package main")
	auxiliar.Validar_email()
}
