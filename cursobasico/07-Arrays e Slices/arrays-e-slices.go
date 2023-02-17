package main

import (
	"fmt"
	"reflect"
)

func main() {
	arrays()
	slices()
	array_interno()
}

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

func arrays() {
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

/*
Slice parece um array dinâmico, porém não é array é uma coleção de dados que tem uma organização
diferente.
Segue a obrigatoriedade de um tipo que irá reger todo o slice.
O Slice permite que a dimensão seja modificada para cima ou para baixo

var slice []int
*/
func slices() {
	fmt.Println("")
	fmt.Println("=========================== Slices ===========================")

	var slice1 []int                                                           //Slice vazio
	var slice2 = []string{"Florianópolis", "São Palo", "Ubatuba", "Guararema"} //Slice inicializado
	array3 := [4]string{"Vagner", "Simone", "Beatriz", "Juliana"}

	fmt.Printf("Valor atual de slice1 = %v\n", slice1)
	slice1 = append(slice1, 5) //Adiciona o valor 5 como elemento do slice.
	slice1 = append(slice1, 6) //Adiciona o valor 6 como elemento do slice.
	fmt.Printf("Após a inclusão de dois itens em slice1 = %v\n", slice1)

	fmt.Printf("Array e slice tem tipos bem distintos.\nEnquanto array3 é do tipo %v, slice2 é do tipo %v \n", reflect.TypeOf(array3), reflect.TypeOf(slice2))

	//É possível criar um slice a partir de um array
	//Porém o slice não é uma cópia e sim ponteiros para os elementos do array

	var array6 = [4]int{1, 1, 1, 4}

	var slice3 = array6[1:3] //Para entender esta sintaxe seria como índice >= 1 && índice < 3, portanto, índices 1 e 2
	fmt.Printf("Conteúdo do slice3 => %v\n", slice3)
	fmt.Printf("Conteúdo original de array6 => %v\n", array6)
	slice3[0] = 2
	slice3[1] = 3
	fmt.Printf("Conteúdo do slice3 após manipulação => %v\n", slice3)
	fmt.Printf("Conteúdo resultante de array6 => %v\n", array6)
}

/*
Todo slice tem como base um array, porém, quando um slice é criado "do nada", este "do nada" cria um array
interno.
*/
func array_interno() {
	fmt.Println("")
	fmt.Println("=========================== Arrays Internos ===========================")
	/*
		A função make irá criar um slice do tipo array de float32, com 10 elementos de tamanho e com capacidade para 32 elementos.
		Slice é um tipo de tamanho total limitado somente pela memória do computador.
		Neste contexto parece estranho definir uma capacidade máxima.
		Este valor na verdade, não é um máximo e sim uma pré-alocação, ou seja, serão alocados previamente espaço em
		memória para um total de 32 elementos.
		Neste caso o slice poderá crescer até 32 elementos sem esforço, quando este limite for rompido, haverá a necessidade
		de mais poder computacional para isto.

		P.S: Se a capacidade não for passada para o make, este utilizará o total de elementos como capacidade
	*/

	slice1 := make([]float32, 10, 15)
	fmt.Printf("slice1 após ser criado pelo comando make: %v\n", slice1)
	fmt.Printf("slice2\ttamanho\t\t: %v\n\tcapacidade\t: %v\n", len(slice1), cap(slice1))

	/*
		Quando a capacidade de um slice é ultrapassada, automaticamente o Go ajusta a capacidade para o dobro.
	*/

	slice2 := make([]string, 2)
	slice2[0] = "Vagner"
	slice2[1] = "Araujo"
	fmt.Printf("slice2 após ser criado pelo comando make: %v\n", slice2)
	fmt.Printf("slice2\ttamanho\t\t: %v\n\tcapacidade\t: %v\n", len(slice2), cap(slice2))

	/*
		Após adicionar 1 elemento o tamanho e a capacidade são ajustados para o dobro do atual
	*/
	slice2 = append(slice2, "Beatriz")

	fmt.Printf("slice2 após inclusçao de mais um elemento: %v\n", slice2)
	fmt.Printf("slice2\ttamanho\t\t: %v\n\tcapacidade\t: %v\n", len(slice2), cap(slice2))

	slice2 = append(slice2, "Aredes")
	slice2 = append(slice2, "Simone")

	fmt.Printf("slice2 após inclusçao de mais um elemento: %v\n", slice2)
	fmt.Printf("slice2\ttamanho\t\t: %v\n\tcapacidade\t: %v\n", len(slice2), cap(slice2))

}
