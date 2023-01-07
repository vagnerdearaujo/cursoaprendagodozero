package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Wait Group permite agrupar e controlar as Go Routines de forma que possamos encerrar o programa
	quando todas as rotinas terminem, independentemente da ordem de execução ou do tempo individual
	de cada uma.

	Em determinados cenários é possível que a execução do programa termine sem que todas as Go Routines
	tenham sido executadas plenamante.
*/

func main() {
	//Cria o grupo de espera
	var waitGroup sync.WaitGroup

	//Informa quantas rotinas serão "vigiadas"
	waitGroup.Add(2)

	//Para o controle do que está sendo aguardado, usa-se uma função anônima
	go func() {
		escrever("Executando de forma concorrente - Membro 1 do WaitGroup")
		waitGroup.Done()
	}()

	go func() {
		escrever("Executando de forma concorrente - Membro 2 do WaitGroup")
		waitGroup.Done()
	}()

	//Valida se todas as rotinas executando de forma concorrente já terminaram.
	waitGroup.Wait() //Sem est comando, nada seria executado.
}

func escrever(texto string) {
	for conta := 1; conta <= 10; conta++ {
		fmt.Printf("%s \t(%v)\n", texto, conta)
		time.Sleep(300 * time.Millisecond)
	}
}
