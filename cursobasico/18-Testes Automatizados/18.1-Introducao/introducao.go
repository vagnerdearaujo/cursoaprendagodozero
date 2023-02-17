package main

import (
	"fmt"
	"testes/enderecos"
)

func main() {
	listaenderecos := []string{"Rua Otelo Augusto Ribeiro", "Praça Clóvis", "Av. Paulista", "Estrada do Lageado"}
	for _, endereco := range listaenderecos {
		tipoEndereco := enderecos.ValidaEndereco(endereco)
		fmt.Printf("O endereço %v é do tipo \"%v\"\n", endereco, tipoEndereco)
	}
}
