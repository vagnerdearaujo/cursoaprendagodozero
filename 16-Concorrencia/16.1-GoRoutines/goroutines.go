package main

import (
	"fmt"
	"time"
)

/*
Concorrência X Paralelismo
Concorrência != Paralelismo
Paralelismo é quando duas ou mais tarefas são executadas ao mesmo tempo, isto somente é possível quando
o computador está equipado com mais de um núcleo, neste caso, as tarefas são distribuídas entre os núcleos
disponíveis.
Tarefas concorrentes são executadas em um ou mais núcleos, porém compartiham o tempo com outras tarefas, ou
seja, o processador executa cada uma das tarefas durante uma fatia de tempo, intercalando a execução entre
cada uma das tarefas.
*/
func main() {
	/*
		Ao iniciar uma chamada com o comando go, o método será executado como um thread, executando,
		portanto de forma concorrente.
		Portanto, ao iniciar uma chamada com Go, informa o compilador para executar o método e não
		esperar o término da ação para executar a próxima.

	*/
	go escrever("Executando de forma concorrente v1") //executa de forma concorrente
	go escrever("Executando de forma concorrente v2") //executa de forma concorrente
	go escrever("Executando de forma concorrente v3") //executa de forma concorrente

	/*
		Atenção: Se todos as chamadas fossem precedidas por Go, nada seria executado, porque após a
		última instrução o programa encerra.
	*/
	escrever("Executando sem o go \t\t")
}

func escrever(texto string) {
	//Cria um loop de 1.000 iterações com uma iteração a cada 300 milisegundos
	for conta := 1; conta < 1000; conta++ {
		fmt.Printf("%s \t(%v)\n", texto, conta)
		time.Sleep(300 * time.Millisecond)
	}
}
