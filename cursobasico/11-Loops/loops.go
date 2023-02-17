package main

import (
	"fmt"
	"time"
)

/*
O Go possui um único comando para repetição o "FOR" com este comando é possível realizar todas
as estruturas de repetição de linguagens como o C#

Um loop infinito pode ser criado usando simplesmente:
	for {}
*/
func main() {
	fmt.Println("=========================== Loops ===========================")

	fmt.Println("=============== While ==================")
	var separador string = ""
	fmt.Print("Construindo separador: ")
	for len(separador) < 15 {
		time.Sleep(100 * time.Millisecond) //Causa uma pausa de 100ms
		separador += "-"
		fmt.Print(".")
	}
	fmt.Println()
	fmt.Println("Separador: " + separador)

	fmt.Println("=============== For Padrão ==================")
	separador = ""
	for i := 0; i < 10; i++ {
		separador += "*"
		fmt.Println(separador)
	}

	fmt.Println("=============== For Each em coleção==================")
	/*
		Esta forma de iterar por coleções usando o range funciona para arrays, slices e maps
	*/

	familia := []string{"Vagner", "Simone", "Beatriz", "Juliana"} //Slices dos Araujos
	for indice, nome := range familia {
		fmt.Println(indice, nome)
	}

	fmt.Println("=============== For Each em string==================")
	frase := "I'm unstoppable, I'm a Porche with no break"
	fmt.Printf("Índice\tASC\tLetra\n")
	for indice, codAsc := range frase {
		fmt.Printf("%v\t%v\t%v\n", indice, codAsc, string(codAsc))
	}

}
