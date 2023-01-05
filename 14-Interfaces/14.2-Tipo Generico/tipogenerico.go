package main

import "fmt"

/*
A interface genérica permite que qualquer coisa declarada com ela se comporte como um variant.
Existem casos muito específicos para seu uso, porém, abre uma porta para todo tipo de gambiarra.
*/

func generica(variant interface{}) {
	fmt.Println(variant)
}

type strteste struct {
	nome  string
	idade uint8
	sexo  string
}

func main() {
	generica("Vagner")
	generica(123.456)
	generica(true)
	generica(strteste{nome: "Vagner", idade: 56, sexo: "M"})

	/*
		Ao declarar uma variável com uma interface genérica, ignora-se completamente checagem de tipos
		do Go. Isto, claro é um risco enorme e abre portas para todo o tipo de gambiarras e más práticas.
		Claro que existem situações específicas em que o uso da interface genérica é imprescindível.
		Ex.: Println que é capaz de receber qualquer tipo de parâmetro.
	*/
	var qqcoisa interface{}
	qqcoisa = "Beatriz"
	fmt.Println(qqcoisa)
	qqcoisa = 8
	fmt.Println(qqcoisa)
	qqcoisa = true
	fmt.Println(qqcoisa)
	qqcoisa = strteste{nome: "Simone", idade: 47, sexo: "F"}
	fmt.Println(qqcoisa)

	//É possível produzir uma anarquia semelhante ao map abaixo
	mapa := map[interface{}]interface{}{
		"nome":  "Vagner",
		"idade": 56,
		1123:    true,
		false:   987.99,
	}

	fmt.Println(mapa)

}
