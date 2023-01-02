package main

import "fmt"

/*
	Não existe herânça em Go, simplesmente porque não existe classes.
	O que existe é um acoplamento de struturas que permite um uso relativamente mais simplificado.
	Por exemplo: é possível criar uma estrutura, usar dento de outra estrutura e ter um acesso
	simplificado às propriedades.
*/

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa      //Neste ponto é onde a magia acontece, é possível declarar uma struct como sendo um pedaço de outra struct.
	curso       string
	instituicao string
}

/*
a struct estudando tem como uma das propriedades uma outra estrutura
Para preencher a estrutura pessoa dentro de estudante é aconselhável criar uma struct pessoa,

	preenche-la e então atribuir à propriedade pessoa dentro de esutdante.

O acesso de leitura se faz de forma natural, ou seja, refere-se às propriedades de pessoal dentro da struct
estudante, como se estas fizessem parte da struct estudante, ou seja:

estudante.nome, estudante.sobrenome, ainda que correto, não há a necessidade de usar fqdn, ou seja:

	estudante.pessoa.nome, estudante.pessoa.altura
*/
func main() {
	fmt.Println("=========================== Herança (Só que não) ===========================")
	var persona = pessoa{nome: "Vagner", sobrenome: "Araujo", idade: 55, altura: 180}
	var estudante1 = estudante{pessoa: persona, curso: "Go Lang do básico ao avançado", instituicao: "Udemy"}
	var estudante2 = estudante{pessoa: pessoa{nome: "Beatriz", sobrenome: "Aredes", idade: 8, altura: 137}, curso: "Fundamental - 3º Ano", instituicao: "COT"}

	fmt.Println(estudante1)
	fmt.Println(estudante2)

	fmt.Printf("O estudante %s, tem %v anos e estuda %s na instituição %s\n", estudante1.nome, estudante1.idade, estudante1.curso, estudante1.instituicao)
	fmt.Printf("O estudante %s, tem %v anos e estuda %s na instituição %s\n", estudante2.nome, estudante2.idade, estudante2.curso, estudante2.instituicao)

}
