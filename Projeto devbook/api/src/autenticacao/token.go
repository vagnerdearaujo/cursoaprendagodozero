package autenticacao

import (
	"api/src/router/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
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

	if _, tokenValido := token.Claims.(jwt.MapClaims); tokenValido && token.Valid {
		return nil
	}
	//fmt.Println(token)
	return errors.New("Token inválido")
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

func TokenIDUsuario(r *http.Request) (uint64, error) {
	tokenCapturado := extrairToken(r)
	if tokenCapturado == "" {
		return 0, errors.New("Token inválido")
	}

	token, erro := jwt.Parse(tokenCapturado, validaSecretKey)
	if erro != nil {
		return 0, errors.New("Token inválido")
	}

	if token.Valid {
		permissoes, tokenValido := token.Claims.(jwt.MapClaims)
		if tokenValido {
			//permissoes é um map[string]interface, portanto não pode ser convertido diretamente para string
			//Por padrão o jwt armazena os tipos numéricos como float, sendo assim, a função Sprintf será utilizada
			//para converter para string a partir de um float.

			usuarioID, erro := strconv.ParseUint(fmt.Sprintf("%.0f",permissoes["usuarioID"]), 10, 64)
			if erro != nil {
				return 0, erro
			}
			return usuarioID, nil
		}
	}

	return 0, errors.New("Usuário não autorizado")
}
