package arrays

import "fmt"

/*
Todo slice tem como base um array, porém, quando um slice é criado "do nada", este "do nada" cria um array
interno.
*/
func Array_interno() {
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
