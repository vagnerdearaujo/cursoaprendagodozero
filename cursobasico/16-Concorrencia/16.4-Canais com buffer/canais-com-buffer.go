package main

import "fmt"

/*
A utilização convencional de um canal é que seja declarado em um método,
e tenha uma mensagem enviada neste mesmo método.
A recepção no entanto, deve ser realizada por outro método.
Por padrão a capacidade de mensagens de um canal é 1.
Ao tentar escrever e ler em um canal no mesmo método, teremos uma mensagem de erro de dead lock.
(fatal error: all goroutines are asleep - deadlock!)
Para bypassar este tipo de problema, pode-se declarar um canal com buffer.
canal := make(chan <tipo>, <tam_buffer>)

A especificação do buffer só funciona neste caso porque o canal é bloqueante.
Quando você especifica um tamanho maior de buffer o canal somente será bloqueado quando o limite
do buffer for atingido.
Por exemplo, se for especificado um buffer de 2, ao tentar enviar uma 3ª mensagem para o canal
teríamos o deadlock
*/
func main() {
	/*
		A execução do código abaixo irá causar o erro de deadlock: fatal error: all goroutines are asleep - deadlock!
		canal := make(chan string)
		canal <- "Este tipo de envio vai causar erro"
		mensagem := <-canal
		fmt.Println(mensagem)
	*/

	canal := make(chan string, 3)
	canal <- "Primeira mensagem"
	canal <- "Segunda mensagem"
	canal <- "Terceira mensagem"

	//canal <- "Quarta mensagem" //Este comando se executado irá levantar a exceção de deadlock

	mensagem := <-canal
	fmt.Println(mensagem)
	mensagem = <-canal
	fmt.Println(mensagem)
	mensagem = <-canal
	fmt.Println(mensagem)

}
