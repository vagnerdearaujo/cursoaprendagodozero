package main

import "fmt"

/*
	As variáveis podem ser declaradas com os tipos explícitos ou implícitos (por inferÊncia)

	Uma vez declarada a variável obrigatoriamente deve ser utilizada ou o Go irá indicar erro de compilação.

	Explícito
		var texto string = "X"
		var numInt int = 3
		var  ()
			text01 string = "Texto 01"
			text02 string = "Texto 02"
			numInt01 int = 3
		)

	Implícito
		texto03 := "Texto 03"
		numInt02 := 25

		Observer que na declaração implícita o operador ":=" deve ser utilizado

	Troca de valores entre variáveis
		texto01, texto02 = texto02,texto01
*/

func main() {
	var (
		text01   string = "Texto 01"
		text02   string = "Texto 02"
		numInt01 int    = 3
	)

	fmt.Println(text01, text02, numInt01)

	text01, text02 = text02, text01
	fmt.Println(text01, text02, numInt01)

	texto03 := "Texto 03"
	numInt02 := 25

	fmt.Println(texto03, numInt02)

}
