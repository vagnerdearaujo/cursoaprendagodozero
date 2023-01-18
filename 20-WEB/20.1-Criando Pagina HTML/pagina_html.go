package main

import (
	"log"
	"net/http"
)

/*
	HTTP: Protocolo de comunicação base da WEB
	Modelo Cliente / Servidor
	Cliente: Faz a Requisição (request)
	Servidor: Envia a resposta (response)
	Rotas: O caminho do recurso a ser "atingido"
		URI: Universal Resource Identificator (Identificador do Recurso)
		Método (Verbo): GET		: Obter informações
						POST	: Enviar informações
						PUT		: Atualização de dados
						DELETE	: Exclusão de dados
*/

func main() {
	/*
		Para colocar um servidor http no ar basta:
		log.Fatal(http.ListenAndServe(":5000",nil))

		Para Indicar uma rota:
		Este comando recebe dois parâmetros.
			1. URI da rota
			2. Função responsável por tratar a requisição
				http.HandleFunc("/home",func(w http.ResponseWriter, r *http.Request) {

				})
				A função pode ser declarada como função anônima ou ser uma função "normal".
				Neste caso a requisição seria: http.HandleFunc("/familia", familia)
				Para este cenário a função familia tem que existir.

				ou seja:
					func familia(w http.ResponseWriter, r *http.Request) {....}

	*/

	//Declara a rota principal
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Opções</h1>"))
		w.Write([]byte("<h3>/home: Para mostrar o Olá Mundo</h3>"))
		w.Write([]byte("<h3>/familia: Para mostrar os integrantes da família Araujo</h3>"))
	})

	//Declara uma rota que mostre Olá Mundo
	http.HandleFunc("/home", hello)

	//Declara uma rota que mostra os integrantes da família Araujo
	http.HandleFunc("/familia", familia)

	//As rotas devem ser declaradas antes de "levantar" o servidor.
	log.Fatal(http.ListenAndServe(":5932", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Olá mundo!!!!</h1>"))
}

func familia(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Família Araujo</h1>"))
	w.Write([]byte("<h2>Vagner de Araujo</h2>"))
	w.Write([]byte("<h2>Simone Aredes Carvalho de Araujo</h2>"))
	w.Write([]byte("<h2>Beatriz Aredes Carvalho de Araujo</h2>"))
	w.Write([]byte("<h2>Juliana Vitória S. de Araujo</h2>"))
}
