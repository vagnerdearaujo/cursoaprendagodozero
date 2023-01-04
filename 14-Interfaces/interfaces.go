package main

import (
	"fmt"
	"math"
)

/*
	A exemplo de outras linguagens como o C# na interface apenas as assinaturas dos métodos são declaradas.
	No entanto a semelhança para aí.
	No Go a interface é implementada de forma implícita, ou seja, não existe uma referência na declaração do
	método que implementa a interface.

	A declaração da interface se assemelha à declaração do struct
*/

type forma interface {
	nomeforma() string
	area() float64
}

type circulo struct {
	raio float64
}

type retangulo struct {
	altura  float64
	largura float64
}

func main() {
	circ := circulo{raio: 27.0}
	ret := retangulo{altura: 14, largura: 17}	
	areaCirculo := circ.area()
	areaRetangulo := ret.area()

	fmt.Println("Obtendo a área via método area da struct")
	fmt.Printf("Área do círculo %0.3f\n", areaCirculo)
	fmt.Printf("Área do retângulo %0.3f\n", areaRetangulo)

	fmt.Println("Obtendo a área via método escreverArea que implementa a interface")
	escreverArea(circ)
	escreverArea(ret)
}

/*
	A exemplo do que acontece com outras linguagens, os métodos que utilizarão a interface, precisam
	explicitar que usam a interface.
	No caso do Go o parâmetro a ser recebido precisa ser uma interface.
*/

func escreverArea(f forma) {
	fmt.Printf("A área calculada da forma %s é: %0.4f\n",f.nomeforma(),f.area())
}

/*
	Para implementar o método area nas formas círculo ou retangulo, basta atachar à struct um método
	com o mesmo nome e assinatura do método encontrado na interface.
*/

/*
Ainda que os métodos tenham o mesmo nome a referência a struct torna o método diferente.

Apenas para exemplificar o método abaixo geraria um erro no Go, pois o nome do método e a struct são as mesmas
ainda que a assinatura do método em sí seja diferente

	func (c circulo) area(vetor int) float64 {
		return math.Pi * math.Pow(c.raio,2)
	}
*/
func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func (c circulo) nomeforma() string {
	return "Círculo"
}

func (r retangulo) area() float64 {
	return r.altura * r.altura
}

func (r retangulo) nomeforma() string {
	return "Retângulo"
}