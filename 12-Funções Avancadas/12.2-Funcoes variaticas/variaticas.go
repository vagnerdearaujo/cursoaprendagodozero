package main

import "fmt"

func main() {
	fmt.Println("=========================== Funções Variáticas ===========================")

	somatoria := somatermos(5, 4, 3, 2, 1, 9, 8, 7, 6)
	fmt.Printf("Total: %v\n", somatoria)

	x := 2
	fmt.Printf("Total  somado = %v\n", somatudoemultiplicaporx(x, 1, 2, 3, 4, 5, 6, 7, 8))
}

/*
	Uma função variática permite recepcionar um número ilimitado de parâmetros do mesmo tipo.
	<nome parâmetro> ... <tipo> cria um slice que recepcionará todos os valores passados à
	função
*/

func somatermos(termos ...int) (total int) {
	total = 0
	for _, numero := range termos {
		total += numero
	}
	return
}

/*
	Um função permite ter um número ilimitado de parâmetros fixos,
	porêm somente um variávl, e o parâmetro variável deve ser o último dos
	parâmetros
*/

func somatudoemultiplicaporx(x int, termos ...int) (total int) {
	total = 0
	for _, numero := range termos {
		total += numero
	}

	total *= x
	return
}
