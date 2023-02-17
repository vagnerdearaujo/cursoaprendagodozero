package main

func main() {
	variavelA, variavelB := 25, 37
	println("Valor de A:", variavelA, " - Valor de B:", variavelB)
	println("Após a inversão dos valores")

	/*
		Ao incluir o "&" à frente da variável o endereço da variável é passada ao invés de seu valor,
		assim o método que receber os endereços conseguirá manipular a variável diretamente na memória
		ignorando assim o escopo no qual a variável foi definida.
	*/
	inverteValores(&variavelA, &variavelB)
	println("Valor de A:", variavelA, " - Valor de B:", variavelB)
}

/*
	Este método é apenas ilustrativo, já que para trocar valores entre variáveis há formas 
	muito mais fáceis e diretas.
*/
func inverteValores(valorA, valorB *int) {
	copiaA := *valorB
	copiaB := *valorA
	*valorA = copiaA
	*valorB = copiaB
}
