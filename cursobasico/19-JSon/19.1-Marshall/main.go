package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

/*
	O Go tem um suporte natural ao formato json, tanto que as declarações de maps assemelham-se
	a uma declaração json.

	Ao declarar uma struct, pode-se incluir os decorators para informar as propriedades dos campos
	Os campos da estrutura a ser serializada devem ser iniciados com uma letra maíuscula.
*/

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	dog := cachorro{"Max", "Vira-Latas", 8}
	fmt.Println(dog)

	cachorroJSON, erro := json.Marshal(dog)
	if erro != nil {
		fmt.Println("Erro na conversão do Json (Cachorro)")
	}

	fmt.Println(bytes.NewBuffer(cachorroJSON))

	//Também é possível utilizar o maps para a conversão.
	/*
		Ao declarar um map o decorator não pode ser utilizado
	*/
	carro := map[string]string{
		"marca":      "Toyota",
		"modelo":     "Yaris",
		"ano":        "2022",
		"fabricacao": "2023",
	}

	fmt.Println(carro)
	carroJSON, erro := json.Marshal(carro)
	if erro != nil {
		fmt.Println("Erro na conversão do Json (Carro)")
	}

	fmt.Println(bytes.NewBuffer(carroJSON))

}
