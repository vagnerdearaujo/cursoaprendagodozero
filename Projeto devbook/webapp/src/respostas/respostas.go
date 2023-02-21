package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// A criação do struct se deve a facilidade de transformar um struct em json
type ErroAPI struct {
	Erro string `json:"erro"`
}

// JSON Escreve na página html cliente
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

// TratarStatusCode Envia para a página o status code e a mensagem de erro
func TratarStatusCode(w http.ResponseWriter, r *http.Response) {
	var erroAPI ErroAPI

	//Decodifica o response body em json e preenche a variável de erro
	json.NewDecoder(r.Body).Decode(&erroAPI)

	//Chama a função que vai escrever no html passando os dados de erro
	JSON(w, r.StatusCode, erroAPI)
}
