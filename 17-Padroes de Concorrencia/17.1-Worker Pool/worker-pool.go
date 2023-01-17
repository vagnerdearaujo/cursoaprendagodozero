package main

import "fmt"

/*
	O padrão worker pools serve para quando se tem um determinado trabalho que se repete
	e pode ser dividio em threads (go routines) de forma a tornar o trabalho mais rápido.

	Ao dividir o trabalho em tarefas (threads) este pode ser concluído mais rapidamente,
	lembrando que a quantidade de tarefas concorrentes vai depender das limitações de
	hardware da tua máquina
*/

func main() {
	fila_fibonnaci, fila_valorescalculados := make(chan int, 45), make(chan int, 45)
	const iteracoes = 45
	const number_of_workers = 4
	for i := 0; i < number_of_workers; i++ {
		go worker(fila_fibonnaci, fila_valorescalculados)
	}

	/*
		ainda que worker tenha sido chamada com a fila vazia, não haverá problemas
		a medida em que a fila contiver valores estes serão processados normalmente
	*/

	for i := 0; i < iteracoes; i++ {
		fila_fibonnaci <- i
	}

	//Ao terminar de alimentar a fila, esta deve ser encerrada.
	close(fila_fibonnaci)

	//Worker irá preencher a fila de resultados
	for i := 0; i < iteracoes; i++ {
		resultado := <-fila_valorescalculados
		fmt.Println(resultado)
	}
}

/*
Ao declarar um canal este pode ser bidirecional, somente leitura ou somente escrita
Canal bidirecional ==> make(chan int)
Canal somente leitura ==> make(<-chan int)
Canal somente escrita ==> make(chan<- int)
*/

/*
Irá chamar o método de calcular o fibonacci, recebendo dois parâmetros:

	um canal de entrada e outro de saída
*/
func worker(fila_fibonnaci <-chan int, fila_valorescalculados chan<- int) {
	for valor := range fila_fibonnaci {
		fila_valorescalculados <- fibonacci(valor)
	}
}

func fibonacci(numero int) int {
	if numero <= 1 {
		return numero
	}
	return fibonacci(numero-2) + fibonacci(numero-1)
}
