package main

/*
	A função init é executada automaticamente como primera coisa do arquivo.
	Pode-se executar a inicialização de variáveis e outras coisas antes que o programa oficialmente
	seja iniciado.

	Apenas para reforçar a função init é executada por arquivo e não por pacote, quer dizer que se um pacote
	tiver 10 arquivos, cada arquivo pode ter seu próprio init que será executado quando o arquivo for invocado.
*/

var eu string
var esteano int
var estemes string

func main() {
	println("Eu:", eu)
	println("Este ano:", esteano)
	println("Este mês:", estemes)
}

func init() {
	println("Executando a função Init e inicializando as variáveis.")
	eu = "Vagner de Araujo"
	esteano = 2023
	estemes = "Janeiro"
}
