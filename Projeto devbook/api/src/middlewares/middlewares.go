package middlewares

import (
	"fmt"
	"net/http"
)

/*
	O middlware é utilizado quando determinados métodos precisam ser executados
	para todas as rotas.
	Ao invés de entrar em cada uma das rotas e configurar a chamada, a utilização
	do meddwlare vai permitir chamar o que se precisa realizar o pŕoprio meddlware
	se encarrega de chamar a função que irá executar o trabalho solicitado pela rota

	É muito comum que no middleware haja um aninhamento de funções, ou seja:
	O middlware recebe como parâmetro a função que deve ser chamada.
	Executa o que precisa ser executado e ao final chama a função que foi injetada
	no middlware via parâmetro
*/

//Autenticar recebe a função original como parâmetro e chama ao final.
//Handler func tem a assinatura das rotas, ou seja, func(w http.ResponseWriter, r *http.Request)

func Autenticar(funcaoOriginal http.HandlerFunc) http.HandlerFunc {
	//Retorna a função que irá lidar com a autenticação

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Validando e autenticando o token .... (sqn)")

		//Chama a função original que foi interceptada pelo middleware
		funcaoOriginal(w, r)
	}
}
