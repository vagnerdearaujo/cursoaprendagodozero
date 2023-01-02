package main

import "fmt"

func main() {
	fmt.Println("=========================== Retorno Nomeado ===========================")

	rsoma, rsub := calculosmatematicos(25, 77)
	fmt.Printf("Soma de X,Y: %v\n", rsoma)
	fmt.Printf("Subtração de Y,X: %v\n", rsub)
}

/*
No retorno nomeado os parâmetros recebem nomes e são retornados automaticamente com o
comando return sem parâmetros

A vantagem deste tipo de abordagem é que as variáveis que normalmente seriam criadas para
o processamento dentro da função.
A grande vantagem é que não precisamos nos preocupar com a sequência de retornos, pois o
próprio compilador se vira
*/
func calculosmatematicos(x, y int) (soma int, subtracao int) {
	soma = x + y
	subtracao = y - x
	return
}
