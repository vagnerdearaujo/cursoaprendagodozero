package main

import "fmt"

/*
Funções anônimas são  blocos de código que se comportam como uma função, porém não
podem ser chamadas explicitamente.
Devem ser executadas no exato momento de sua criação;

Para ser sincero não vi utilidade neste tipo de função
*/
func main() {
	total := func(valores ...int) int {
		soma := 0
		for _, valor := range valores {
			soma += valor
		}
		return soma
	}(2, 4, 6, 8, 10, 12, 14, 16, 18)

	fmt.Println(total)

}
