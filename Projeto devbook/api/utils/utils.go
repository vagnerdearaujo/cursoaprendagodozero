package utils

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

func GeraSecretKey() string {
	chave := make([]byte, 64)

	//Popula o slice de bytes
	if _, erro := rand.Read(chave); erro != nil {
		log.Fatal(erro)
	}

	return base64.StdEncoding.EncodeToString(chave)
}
