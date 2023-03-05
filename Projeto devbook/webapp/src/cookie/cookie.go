package cookie

import (
	"net/http"
	"time"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

/*
	Tanto o pacote mux quanto o securecookie tem os mesmos criadores.
	Para maiores informações acesse o https://www.gorillatoolkit.org/
	Acessado em 22/02/2023
*/

var scookie *securecookie.SecureCookie

// ConfigurarCookie Cria um cookie com as informações do arquivo de configuração.
func ConfigurarCookie() {
	scookie = securecookie.New(config.HashKey, config.BlockKey)
	/*
		BlockKey is optional, used to encrypt values.
		Create it using GenerateRandomKey().
		The key length must correspond to the block size of the encryption algorithm.
		For AES, used by default, valid lengths are 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.

		De acordo com a ajuda da função New as chaves precisam ter o tamanho de 16,24 ou 32 caracteres.
	*/
}

// ArmazenaCookie armazena no browser o cookie com i ID e Token do usuário
func ArmazenaCookie(w http.ResponseWriter, ID, token string) error {
	dados := map[string]string{
		"id":    ID,
		"token": token,
	}

	//Codificar os dados para geração do cookie
	dadosCodificados, erro := scookie.Encode(config.CookieName, dados) //Nome do cookie, Dados geradores do cookie
	if erro != nil {
		return erro
	}

	/*
		Para se certificar que o cookie foi armazenado, verifique no browser após o login
		Modo de desenvolvedor normalmente (F12).
		Applications->Cookies
	*/
	//Armazenar o cookie na página
	http.SetCookie(w, &http.Cookie{
		Name:     config.CookieName, //Nome do cookie
		Value:    dadosCodificados,  //Dados do cookie
		Path:     "/",               //Informa que deve funcionar em todo o site
		HttpOnly: true,              //Ajuda a mitigar o risco do cookie ser acessado pelo lado do cliente
	})
	return nil
}

// CarregarCookie Carrega as informações armazenadas no cookie se houver.
func CarregarCookie(r *http.Request) (map[string]string, error) {
	//Carrega os dados do Cookie que estão na requisição
	cookie, erro := r.Cookie(config.CookieName)
	if erro != nil {
		return nil, erro
	}

	//Cria um map vazio para receber os dados do cookie.
	valoresCookie := make(map[string]string)

	//Decodifica os dados do cookie
	if erro = scookie.Decode(config.CookieName, cookie.Value, &valoresCookie); erro != nil {
		return nil, erro
	}

	return valoresCookie, nil
}

// LimparCookie: Limpa os valores do cookie e o torna expirado para forçar um novo login
func LimparCookie(w http.ResponseWriter) {
	//Armazenar o cookie na página
	http.SetCookie(w, &http.Cookie{
		Name:     config.CookieName, //Nome do cookie
		Value:    "",                //Passa um valor em branco para forçar cookie inválido
		Path:     "/",               //Informa que deve funcionar em todo o site
		HttpOnly: true,              //Ajuda a mitigar o risco do cookie ser acessado pelo lado do cliente
		Expires:  time.Unix(0, 0),   //Informa que o cookie já está expirado
	})
}
