package main

import (
	"fmt"
	"time"
)

/*
	Canal é um mecanismo que o Go implementa para enviar e receber dados.
	Diferentemente de variáveis os canais podem ser utilizados para a comunicação entre threads de forma
	segura.
*/

func main() {
	//A exemplo de qualquer entidade no Go o canal deve ser declarado com o tipo de dado que irá trafegar
	canal := make(chan string)
	go escrever("Mensagem 1", canal)

	/*
		Aguarda até que o canal receba um dado.
		Um cuidado que se deve ter tem relação com o dead-lock.
		Esta condição ocorre quando um canal um método que deveria enviar dados para o canal é terminado.
		Neste cenário o canal para de receber dados e não há mais nenhum outro método para fazê-lo.

		Além de receber o dado do canal, também é possível receber seu estado (aberto/fechado)
	*/

	aberto := true
	mensagem := ""
	for aberto {
		mensagem, aberto = <-canal
		fmt.Println(mensagem)
	}

}

func escrever(texto string, canal chan string) {
	for conta := 1; conta <= 10; conta++ {
		/*	O formato de envio de dados para o canal é diferente da atribuição para uma variável,
			ainda assim é bastante intuitivo.
		*/

		canal <- fmt.Sprintf("%v (%v)\n", texto, conta)
		//fmt.Printf("%v (%v)\n", texto, conta)
		time.Sleep(300 * time.Millisecond)
	}

	//Informa que o canal está fechado, assim quem estiver "escutando" o canal, saberá que nenhuma outra mensagem
	//será enviado através do canal
	close(canal)
}
