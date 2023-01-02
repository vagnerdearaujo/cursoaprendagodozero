package main

import "fmt"

/*
	Ponteiros não armazenam valores, ponteiros armazenam o endereço de outras variáveis.
	Este mecanismo é utilizado para permitir que uma variável possa ser acessada fora de seu escopo.
	Em algumas linguagens é possível passar uma variável em referência (ponteiro) com o uso de alguma palavra reservada.
	Em Go isto deve ser feito explícitamente passando o ponteiro de uma variável para que esta possa ser modificada
	fora de seu contexto original.
*/

func main() {
	fmt.Println("Ponteiros")
	var variavel1 int = 10
	var variavel2 int = variavel1
	fmt.Printf("variavel1 = %v , variavel2 = %v\n", variavel1, variavel2)

	/*
		A declaração do ponteiro é semelhante à declaração de uma variável.
		Importante notar que ainda que um ponteiro seja o endereço de memória de uma outra variável,
		ainda assim precisa ser declarado como o mesmo tipo da variável a ser apontada
		Como qualquer variável os ponteiros podem ser recebidos via parâmetros ou serem retornados por funções
	*/

	//Pode-se declarar o ponteiro ou declarar já atribuindo o "alvo"
	//É preciso informa explicitamente que o que está sendo passado para o ponteiro é seu endereço
	//Para sinalizar isto precede-se a variável alvo com & (e comercial)
	var ponteiro1 *int
	// O ponteiro conterá nil se declarado e não tiver um valor atribuído.

	ponteiro1 = &variavel1

	var ponteiro2 *int = &variavel2

	fmt.Printf("Endereço 1 = %v, Endereço 2 = %v \n", ponteiro1, ponteiro2)

	/*
		Para interagir com os valores apontados pelos endereços de memória dos ponteiros, é preciso realizar
		a "desreferenciação", ou seja, uso o ponteiro como se fosse a própria variável.
		Para isto o ponteiro deve vir precedido por um "*" (asterísco)
		Um ponteiro precedido por um * é a própria variável sendo acessada não por seu nome e sim por seu
		endereço de memória.
	*/

	*ponteiro1 = 25 //Troca o valor da variável 1
	*ponteiro2 = 50 //Troca o valor da variável 2
	fmt.Printf("variavel1 = %v , variavel2 = %v\n", variavel1, variavel2)
}
