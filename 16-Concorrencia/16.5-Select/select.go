package main

import (
	"fmt"
	"time"
)

/*
O comando select é semelhante ao switch, porém é utilizado exclusivamente com canais.
*/

func main() {
	canal1, canal2 := make(chan string), make(chan string)
	iteracoes := 100

	/*
		No cenário sem select a leitura do 1º canal precisa esperar a leitura 2º canal
		para que os resultados sejam apresentados.

		O correto é que para cada mensagem 2 tivessemos 4x a mensagem 1, pois o tempo de espera é
		1/2s contra 2s.
	*/
	for i := 0; i < iteracoes; i++ {
		enviamensagem_canal(canal1, canal2, iteracoes)
		mensagem1 := <-canal1
		mensagem2 := <-canal2
		fmt.Println(mensagem1)
		fmt.Println(mensagem2)
	}

	/*
		No cenário com select os canais são processados à medida em que possuem
		valores para serem lidos, permitindo que a leitura ocorra sem bloqueios
	*/

	for i := 0; i < iteracoes; i++ {
		select {
		case mensagem1 := <-canal1:
			fmt.Println(mensagem1)
		case mensagem2 := <-canal2:
			fmt.Println(mensagem2)
		}
	}
}

func enviamensagem_canal(canal1, canal2 chan string, iteracoes int) {
	go func(vezes int) {
		for i := 0; i < vezes; i++ {
			time.Sleep(time.Microsecond * 500) //Pausa de 1/2s
			canal1 <- "Canal 1"
		}

	}(iteracoes)

	go func(vezes int) {
		for i := 0; i < vezes; i++ {
			time.Sleep(time.Microsecond * 2) //Pausa de 2s
			canal2 <- "Canal 2"
		}

	}(iteracoes)
}
