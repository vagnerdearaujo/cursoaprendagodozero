package main

import "fmt"

/*
	O struct equivale ao register do Delphi.
	É uma estrutura ou coleção de tipos, porém, sem métodos.
*/

type usuario struct {
	nome  string
	idade uint8
}

/*
- Formas de inicializar uma variável de um tipo strunct
- var v1 usuario
- v1.nome = "Vagner"
- v1.idade = 55

- var v2 usuario = usuario{"Vagner",55}
-
- v3 := usuario{"Simone",47}

- //Quando houver a necessidade de omitir campos da struct
- v4 := usuario{nome: "Juliana"}
*/
func main() {
	var usr1 usuario = usuario{"Vagner", 55}

	var usr2 usuario
	usr2.nome = "Simone"
	usr2.idade = 47

	usr3 := usuario{"Beatriz", 9}
	usr4 := usuario{nome: "Juliana"}
	fmt.Println(usr1)
	fmt.Println(usr2)
	fmt.Println(usr3)
	fmt.Println(usr4)

}
