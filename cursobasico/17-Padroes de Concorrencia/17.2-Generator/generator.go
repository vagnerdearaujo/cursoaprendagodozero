package main

import (
	"fmt"
	"time"
)

/*
	Basicamente o padrão generator encapsula uma chamada a uma Go routine e devolve um canal
*/

// Este struct não estava no código original eu o incluí apenas para entender como seria passar
// um canal do tipo struct
type magic_chanel struct {
	vezes int
	texto string
}

func main() {
	/*
		Note que a função main, não vai ser responsável por chamar uma go routine, ou mesmo criar
		o canal via make
	*/

	canal := escrever("Chamada a generator")
	for i := 0; i < 100; i++ {
		retorno := <-canal
		fmt.Printf("%s (%3d)\n", retorno.texto, retorno.vezes)
	}
}

func escrever(texto string) <-chan magic_chanel {
	voltas := 0
	canal := make(chan magic_chanel)

	//Chamada a uma go routine, que tanto pode ser uma função anônima, quanto uma outra função qualquer
	go func() {
		for {
			voltas++
			canal <- magic_chanel{voltas, texto}
			time.Sleep(time.Millisecond * 200)
		}
	}()

	return canal
}
