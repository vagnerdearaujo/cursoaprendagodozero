package requisicoes

import (
	"io"
	"net/http"
	"webapp/src/cookie"
)

// FazerRequisicaoComAutenticacao Faz uma requisição http para a API enviando o método e os dados do token.
func FazerRequisicaoComAutenticacao(r *http.Request, metodo, url string, dados io.Reader) (*http.Response, error) {
	/*
		Parâmetros
			r: Requisição original da página para o WebApp. É desta requisição que é extraído o cookie
			metodo: Delete, Get, Post, Put
			url: URL da API que irá processar a requisição
			dados: Quando houver parâmetros, será passado neste parâmetros

		Retorno
			*http.Response: Retorna a resposta da API
			error: auto-explicativo
	*/

	//Cria uma nova requisição para chamar a API.
	request, erro := http.NewRequest(metodo, url, dados)
	if erro != nil {
		return nil, erro
	}

	//Lê os dados do cookie.
	//O erro será ignorado pois o cookie já foi validado na chamada do middleware
	cookie, _ := cookie.CarregarCookie(r)

	//Injeta na requisição o token
	//O token é do tipo Bearer, portanto esta informação também precisa ser passada na injeção do token
	request.Header.Add("Authorization", "Bearer "+cookie["token"])

	//Cria o client que fará a requisição http para a API
	client := &http.Client{}

	//Executa a chamada à API via client
	response, erro := client.Do(request)
	if erro != nil {
		return nil, erro
	}

	return response, nil
}
