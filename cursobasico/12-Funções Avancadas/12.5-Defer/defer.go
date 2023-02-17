package main

import "fmt"

/*
Defer não é um tipo de função e sim o uso para uma.
Defer faz com que um método tenha sua execução adiada.
Na prática, signfica que o método será envocado somente quando o bloo de código onde
foi chamado terminar.

O defer pode ser muito útil para ações que o método terá que executar independentemente do seu
término, por exemplo encerrar a conexão com o banco de dados.

Neste cenário, não importa se os procedimento foram ou não executados com sucesso, a conexão será
encerrada.
*/
func main() {
	fmt.Println("=========================== Defer ===========================")

	teste01()
	teste02()

}

func teste01() {
	println("Teste 01")
	/*
		Neste bloco os metodo1 e metodo2 são chamados e executados na mesma ordem
	*/
	println("Execução sem o Defer")
	metodo1()
	metodo2()

	/*
		Neste bloco os metodo1, metodo2 e metodo3 são chamados em ordem, porém o metodo1 é o
		último a ser exedcutado porque o defer adiou sua execução até o último momento.
	*/
	println("Execução com o Defer")
	defer metodo1()
	metodo2()
	metodo3()

	media := calculamediasimples(13.5, 17.9)

	fmt.Println(media)
}

func teste02() {
	println("==============================================================================")
	println("Teste 02")
	/*
		Neste bloco os metodo1 e metodo2 são chamados e executados na mesma ordem
	*/
	println("Execução sem o Defer")
	metodo1()
	metodo2()

	/*
		Neste bloco os metodo1, metodo2 e metodo3 são chamados em ordem, porém o metodo1 é o
		último a ser exedcutado porque o defer adiou sua execução até o último momento.
	*/
	println("Execução com o Defer chamando os métodos na ordem natural, todos com defer")
	defer metodo1()
	defer metodo2()
	defer metodo3()

	/*
	  O defer executa os métodos no esquema lifo.
	  faz sentido por se tratar de uma pilha de execução.
	*/

	media := calculamediasimples(13.5, 17.9)

	fmt.Println(media)
}

func metodo1() {
	println("Chamando o método 1 ")
}

func metodo2() {
	println("Chamando o método 2 ")
}

func metodo3() {
	println("Chamando o método 3 ")
}

func calculamediasimples(valor1, valor2 float64) (media float64) {
	/*
		Não havia a menor necessidade de criar uma função anônima neste ponto e adiar sua execução.
		O que ficou provado é que apesar do código ter sido executado por último guardou o valor
		original para a variável media.

		Portanto para ter o valor atual da variável media, transformei o parâmetro em ponteiro,
		neste cenário o valor correto foi apresentado.

		Após implementar a função anônima neste método, finalmente vi um uso útil para ela.
		Com uma função anônima é possível ter um bloco de código que poderá ser executado posteriormente
		via defer.

		Em uma situação de encerramento de comunicação ou fechamento de banco de dados e gravação
		de log, este esquema de função anônima cairia muito bem.
	*/

	//inicializar a variável média com 33.3 apenas para provar que o defer pega o estado da variável
	//mesmo que seja executado depois.
	media = 33.3 //Este valor vai ser apresentado independentemente do resultado do cálculo
	defer func(ponteiroparamedia *float64, valormedia float64) {

		fmt.Printf("Média calculada (ponteiro) %4.4f\n", *ponteiroparamedia)
		fmt.Printf("Média calculada (valor) %4.4f\n", valormedia)
	}(&media, media)

	media = (valor1 + valor2) / 2
	return
}
