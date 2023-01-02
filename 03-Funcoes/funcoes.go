package main

/*
	Em Go como em C, não existe um comando para a criação de procedures, todos os métodos são funções (func).
	Uma procedure é um método sem retorno, em Go, basta não definir o tipo de retorno, para se ter uma "procedure".
	As funções seguem basicamente o mesmo esquema das declarações Delphi, ou seja:
	func <nome da função ([parâmetros]) <tipo> ou
	func <nome da função ([parâmetros]) (<tipo>,<tipo>,<tipo..n>)

	Se uma função tiver mais de um retorno, os valores devem ser retornados na mesma ordem  e tipo da declaração.
	Quando um retorno não interessar, pode-se substituí-lo por um "_" (underscore)

	Os métodos utilizados na func main podem ser declarados antes ou depois dela.

	Particularmente eu prefiro declarar as funções usadas no main depois dela e não antes.
*/

import "fmt"

func main() {
	fmt.Println("=========================== Funções ===========================")
	var (
		val1           int = 47                 //Declaração explícita
		val2           int = 23                 //Declaração explícita
		soma               = soma(val1, val2)   //Declaração por inferência
		divisao, resto     = divide(val1, val2) //Declaração por inferência
	)
	fmt.Printf("Valores: X = %v  - Y = %v\n", val1, val2)
	fmt.Printf("Soma: %v\n", soma)
	fmt.Printf("Divisão: %3.4f\n", divisao)
	fmt.Printf("Resto: %3.4f\n", resto)
}

func soma(x, y int) int {
	return x + y
}

func divide(x, y int) (float64, float64) {
	resultado := float64(x / y) //Força para o tipo de retorno, se isto não for feito, não existirá a conversão implícita de int para float.
	resto := float64(x % y)
	return resultado, resto
}
