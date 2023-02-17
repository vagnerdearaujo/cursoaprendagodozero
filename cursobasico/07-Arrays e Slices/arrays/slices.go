package arrays

import (
	"fmt"
	"reflect"
)

/*
Slice parece um array dinâmico, porém não é array é uma coleção de dados que tem uma organização
diferente.
Segue a obrigatoriedade de um tipo que irá reger todo o slice.
O Slice permite que a dimensão seja modificada para cima ou para baixo

var slice []int
*/
func Slices() {
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
