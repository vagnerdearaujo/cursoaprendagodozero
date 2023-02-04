package autenticacao

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CriarToken(usuarioID uint64) (string, error) {
	//Mapeia as permissões do usuário
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(3 * time.Hour).Unix() //Expira em 3h
	permissoes["usuarioID"] = usuarioID

	//Gera o token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString([]byte("senha"))
}
