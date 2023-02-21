package config

var (
	//URL da API
	APIURL = ""

	//Porta da API
	Porta = 0

	//HaskKey é utilizado para autenticar o cookie
	HashKey []byte

	//BlockKey é utilizado para criptografar o cookie
	BlockKey []byte
)

func CarregarVariaveisAmbiente() {
	if erro := godotenv.Load();erro != nil {
		log.Fatal(erro)
	}
}