package arrays

import "fmt"

/*
	Arrays = vetor.
	A exemplo das linguagens fortemente tipadas todos os elementos de um array são obrigatoriamente do mesmo tipo
	base do array, o que quer dizer que um array do tipo int, somente conterá elementos deste tipo.
	É Obrigatorio definir a dimensão do array

	a declaração básica de um array
	var <nome array>[<dimensão>] <tipo>
	var meu_array[7] int

	A exemplo da maioria das linguagens o primeiro elemento do array tem índice 0
*/

func Arrays() {
	fmt.Println("=========================== Arrays ===========================")
	var array1 [5]int
	var array2 = [5]int{1, 2, 3, 4, 5}
	array3 := [4]string{"Vagner", "Simone", "Beatriz", "Juliana"}

	fmt.Println(array1, array2, array3)

	//Ainda que a declaração do array exija a definição da dimensão, é possível criar um array e deixar o Go
	//definir o tamanho da estrutura com base na quantidade de itens que o inicializaram

	var array4 = [...]int{55, 47, 8, 21, 74, 82}
	fmt.Println(array4)

	var array5 = array4

	//Quando se copia um array é realmente uma cópia e não uma passagem por referência de forma implícita
	array5[0] = 56

	fmt.Println(array4, array5)

}
