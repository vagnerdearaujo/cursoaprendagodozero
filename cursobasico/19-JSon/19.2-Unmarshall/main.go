package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

/*
	O método unmarshal popula uma variável de struct a partir de um json
*/

func main() {
	cachorroJSON := `{"nome":"Max","raca":"Vira-Latas","idade":8}`

	var dog cachorro

	/*
		A função unmarshal recebe os dados json em um slice de bytes e a variável do tipo da estrutura passada por referência
	*/

	if erro := json.Unmarshal([]byte(cachorroJSON), &dog); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(dog)

	/*
		O mesmo pode ser feito para um map
	*/
	carroJSON := `{"ano":"2022","fabricacao":"2023","marca":"Toyota","modelo":"Yaris"}`
	var carro map[string]string

	json.Unmarshal([]byte(carroJSON), &carro)

	fmt.Println(carro)
}
