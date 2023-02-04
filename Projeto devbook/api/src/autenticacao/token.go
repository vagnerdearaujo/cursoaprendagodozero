package autenticacao

import (
	"api/src/router/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
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
	return token.SignedString([]byte(config.API_SecretKey))
}

func ValidarToken(r *http.Request) error {
	tokenCapturado := extrairToken(r)
	if tokenCapturado == "" {
		return errors.New("Token inválido")
	}

	//Por uma questão de segurança a função parse obriga que o token seja validado
	//para garantir uma consistência e segurança.
	//Claro que a função de verificação poderia simplesmente devolver api_secretkey e
	//um nil para o erro, porém, não seria uma boa prática.
	token, erro := jwt.Parse(tokenCapturado, validaSecretKey)
	if erro != nil {
		return erro
	}
	fmt.Println(token)

	return nil
}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	respostas := strings.Split(token, " ")
	//Por padrão o header conterá Bearer <<token>>
	//Se houver duas palavras, está no padrão, porém não dá para garantir que esteja tudo certo
	if len(respostas) == 2 {
		return respostas[1]
	}
	return ""
}

func validaSecretKey(token *jwt.Token) (interface{}, error) {
	if _, assinaturavalida := token.Method.(*jwt.SigningMethodHMAC); !assinaturavalida {
		return nil, fmt.Errorf("Método de assinatura inesperado %v", token.Header["alg"])
	}
	return config.API_SecretKey, nil
}
