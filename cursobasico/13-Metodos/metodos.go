package main

/*
	Ainda que visualmente métodos e funções sejam exatamente a mesma coisa, conceitualmente não.
	Em linguagens orientadas à objetos os métodos estão intimamente ligados à classes.
	Em Go não existem classes, portanto, métodos estão atachados à structs e interfaces.
*/

type usuario struct {
	nome  string
	idade uint8
}

func main() {
	user1 := usuario{nome: "Vagner de Araujo", idade: 56}
	var user2 usuario
	user2.inicializa("Simone Aredes", 47)

	user3 := usuario{nome: "Beatriz Aredes", idade: 8}

	user1.dados()
	user2.dados()
	if user3.maiordeIdade() {
		println("O usuário ", user3.nome, " é maior de idade")
	} else {
		println("O usuário ", user3.nome, " é menor de idade")
	}
}

/*
Para atachar um método a uma estrutura é preciso referenciar este estrutura antes mesmo do nome do método.
*/
func (u usuario) dados() {
	println("Nome: ", u.nome, "  -  Idade:", u.idade)
}

/*
Se ao invés de passar a estrutura, for passado o ponteiro da estrutura os métodos passarão a ter o poder
de alterar os dados dentro da estrutura passada
*/
func (u *usuario) inicializa(nome string, idade uint8) {
	u.nome = nome
	u.idade = idade
}

func (u usuario) maiordeIdade() bool {
	return u.idade >= 18
}
