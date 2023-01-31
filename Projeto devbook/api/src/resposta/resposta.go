package resposta

import (
	"encoding/json"
	"log"
	"net/http"
)

// Json recebe a inteface da página, código de erro e os dados a serem convertidos como interface vazia.
func JSon(w http.ResponseWriter, statuscode int, dados interface{}) {
	w.WriteHeader(statuscode)
	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}

}

// Recebe o erro e o status e chama a função de conversão para JSon
// Uma vez que Json aceita uma interface vazia, significa que qualquer dado pode ser passado
func Erro(w http.ResponseWriter, statuscode int, erro error) {
	JSon(w, statuscode, struct {
		Erro string `json:"erro"`
	}{Erro: erro.Error()})
}
