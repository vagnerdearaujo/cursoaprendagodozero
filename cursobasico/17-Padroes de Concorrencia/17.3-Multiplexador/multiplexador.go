package main

import (
	"fmt"
	"time"
)

/*
A ideia deste padrão é juntar dois ou mais canais em apenas um
Isto é feito através de uma função com um select que irá retornar o valor dos canais em um canal de saída
*/
func main() {
	canalMultiplexado := multiplexar(escrever("Canal 1"), escrever2("Canal 2"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canalMultiplexado)

	}

}

func multiplexar(canalDeEntrada1, canalDeEntrada2 <-chan string) <-chan string {
	canalRetorno := make(chan string)
	go func() {
		for {
			select {
			case retorno := <-canalDeEntrada1:
				canalRetorno <- retorno
			case retorno := <-canalDeEntrada2:
				canalRetorno <- retorno
			}
		}
	}()
	return canalRetorno
}

func escrever(texto string) <-chan string {
	voltas := 0
	canal := make(chan string)

	//Chamada a uma go routine, que tanto pode ser uma função anônima, quanto uma outra função qualquer
	go func() {
		for {
			voltas++
			canal <- fmt.Sprintf("%s (%3d)", texto, voltas)
			time.Sleep(time.Millisecond * 200)
		}
	}()

	return canal
}

func escrever2(texto string) <-chan string {
	voltas := 0
	canal2 := make(chan string)

	//Chamada a uma go routine, que tanto pode ser uma função anônima, quanto uma outra função qualquer
	go func() {
		for {
			voltas++
			canal2 <- fmt.Sprintf("Escrever 2: %s (%3d)", texto, voltas)
			time.Sleep(time.Millisecond * 200)
		}
	}()

	return canal2
}
