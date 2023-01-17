package enderecos

import "strings"

// ValidaEndereco verifica se o endereço é iniciado por um dos tipos esperados
func ValidaEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "praça", "rodovia"}
	tipoEncontrado := strings.Split(endereco, " ")[0] //Cria um slice de endereço separado por espaços e pega somente a primeira posição

	//Garante que o resultado seja minúsculo
	tipoEncontrado = strings.ToLower(tipoEncontrado)

	tipoRetorno := "Tipo Inválido"

	for _, tipo := range tiposValidos {
		if tipo == tipoEncontrado {
			tipoRetorno = tipo
		}
	}

	return tipoRetorno
}
