package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/cookie"
)

/*
	Para permitir o encadeamento de funções, as funções terão como parâmetro de entrada um handlerfunc e também
	retornarão um handlerfunc.
	Isto permite criar  um "gancho", pois me permite receber a função a ser executada, executar algo e passar
	a função que deverá ser chamada a seguir.

	Com esta técnica é possível executar ou não a função contida no parâmetro proximaFuncao e enviar ou não
	a função a ser executada a seguir.
*/

// Logger Envia para o terminal informações básicas da requisição
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	//Retorna o corpo da execução + a chamada à próxima função.
	return func(w http.ResponseWriter, r *http.Request) {
		//Loga as informações básicas da requisição.
		log.Printf("\n%s %s %s", r.Method, r.RequestURI, r.Host)

		//Inclui na construção da função a chamada à função que deve ser executada a seguir.
		proximaFuncao(w, r)
	}
}

// ValidaInformacaoCookie Verifica se as informações do cookie estão presentes, não valida a informação.
func ValidaInformacaoCookie(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		valoresCookie, erro := cookie.CarregarCookie(r)
		fmt.Println(valoresCookie, erro)
		proximaFuncao(w, r)

	}
}
