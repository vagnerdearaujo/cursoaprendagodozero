package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, statuscode int, dados interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statuscode)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}
