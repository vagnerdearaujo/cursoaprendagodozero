package utils

import "net/http"

func EscreveNaPagina(w http.ResponseWriter, mensagem string) {
	w.Write([]byte(mensagem))
}
