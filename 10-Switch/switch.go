package main

import "fmt"

/*
	Diferentemente do switch do C# o switch em go não usa a clausula break.
	Também é possível omitir a variável de avaliação, colocando a condição diretamente na linha do case
	ex.: case variavel == 4:

	Quando uma ou mais valores satisfazem uma determinada condição, estes podem ser separados por vírgulas
	ex.: case 8, 9, 10:

	Para casos muito excepcionais existe a clausula: fallthrough
		 Esta cláusula "cai para a próxima" cláusula.
		 Se uma determinada condição foi atendida ela é executada e se houver o fallthrough a próxima condição
		 é executada como verdadeira.
*/

func main() {
	fmt.Println("=========================== Switch ===========================")
	fmt.Println(diaDaSemana(2))
	fmt.Println(diaDaSemana(5))
	fmt.Println(diaDaSemana(8))

}

func diaDaSemana(dia int) string {
	var diasemana string
	switch dia {
	case 1:
		diasemana = "Domingo"
	case 2:
		diasemana = "Segunda"
	case 3:
		diasemana = "Terça"
	case 4:
		diasemana = "Quarta"
	case 5:
		diasemana = "Quinta"
	case 6:
		diasemana = "Sexta"
	case 7:
		diasemana = "Sábado"
	case 8, 9, 10:
		diasemana = "Maluco" //Esta cláusula é atendida e executada, portanto neste ponto diasemana = "Maluco"
		fallthrough          //O fallthrough fará com que o resultado seja a clausula default
	default:
		diasemana = "Inválido"
	}
	return diasemana

}
