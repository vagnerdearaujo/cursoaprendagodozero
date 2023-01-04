package main

/*
	Closure são funções que referenciam variáveis que estão fora de seu corpo

	Ainda que a variável texto tenha sido declarada com o mesmo nome da variável dentro da função funcClosure
	o que será impresso ao chamar a variável funcaoNova é o conteúdo da variável do escopo da função funcClosure.
*/
func main() {
	texto := "Dentro da função main"
	texto += "" //Apenas para evitar o erro no go.

	funcaoNova := funcClosure()
	funcaoNova()
}

/*
	Lembrando que inclusive funções são tipos de dados em go, significa que é possível passar e receber
	funções como parâmetros, bem como inicializar uma variável como uma função.
*/

func funcClosure() func() {
	texto := "Dentro da função closure()"

	funcao := func() {
		println(texto)
	}

	return funcao
}
