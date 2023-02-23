package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
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
		if config.Logger {
			log.Printf("\n%s %s %s", r.Method, r.RequestURI, r.Host)
		}

		//Inclui na construção da função a chamada à função que deve ser executada a seguir.
		proximaFuncao(w, r)
	}
}

// ValidaInformacaoCookie Verifica se as informações do cookie estão presentes, não valida a informação.
func ValidaInformacaoCookie(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		valoresCookie, erro := cookie.CarregarCookie(r)
		if config.Logger {
			fmt.Println(valoresCookie, erro)
		}

		//Se houver erro no obtenção do cookie, a página deverá ser redirecionada para o login.
		//a função redirect exige um statuscode que deve ser classe 300, neste caso será usado
		//302 "Found" - Foi encontrado, porém redirecionado
		if erro != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}
		proximaFuncao(w, r)
	}
}
