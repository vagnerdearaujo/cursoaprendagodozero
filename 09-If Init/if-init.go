package main

import "fmt"

func main() {
	fmt.Println("=========================== If Init ===========================")

	/*
		if - init
		esta estrutura permite criar uma variável válida somente no escopo do if.
		este tipo de variável temporária costuma ser criada para situações muito específicas e
		quando criada dentro do método fica "morta" para o restante do método.
		Usada também para evitar a repetição de condições baseadas em cálculos que em algumas situações
		precisam ser repetidos diversas vezes na condição.
	*/

	var numero int = 2
	if outronumero := (numero * 4); outronumero >= 7 {
		fmt.Println("Número maior que 7")
	} else {
		fmt.Println("Número menor que 7")
	}

	/*
		A variável outronumero existe somente no contexto do if
	*/

}
