package main

func main() {
	fmt.Println("=========================== Funções Recursivas ===========================")

	fibo := fibonacci(7)
	println(fibo)
}

/*
	Funções recursivas são muito traiçoeiras pois carregam três problemas complicados
	1. Precisam de uma condição de saída para evitar um loop infinito
	2. Dependendo da quantidade de iterações o tempo pode ser muito alto
	3. Podem gerar uma exceção de stack overflow
*/
func fibonacci(valor uint) uint {
	if valor <= 1 {
		return valor
	}
	return fibonacci(valor-2) + fibonacci(valor-1)
}
