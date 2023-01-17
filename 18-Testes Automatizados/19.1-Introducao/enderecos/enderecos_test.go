// Teste de Unidade
package enderecos

import (
	"testing"
)

/*
Em GO as funções de testes devem seguir algumas normas:
1. O arquivo fonte deve obrigatoriamente terminar com "_test.go"
2. O nome do método começa obrigatoriamente com a palavra Test (com T maíusculo) e o nome
da função a ser testada.
*/

/*
Estrutura de cenários para permitir testar todos os tipos válidos dentro de ValidaEnderecos
*/
type cenarioteste struct {
	endereco     string
	tipoesperado string
}

// "rua", "avenida", "estrada", "praça", "rodovia"}
func TestValidaEndereco(t *testing.T) {
	cenarios := []cenarioteste{
		{"Rua da Glória", "rua"},
		{"Avenida Paulista", "avenida"},
		{"avenida rebouças", "avenida"},
		{"Praça da Árvore", "praça"},
		{"PRAÇA BEVILÁQUA", "praça"},
		{"Estrada M'Boi Mirim", "estrada"},
		{"Rodovia dos Bandeirantes", "rodovia"},
		{"Sei lá onde", "Tipo Inválido"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenarios {
		tipoRecebido := ValidaEndereco(cenario.endereco)
		if cenario.tipoesperado != tipoRecebido {
			t.Errorf("Tipo esperado \"%v\" é diferente do tipo recebido \"%v\"", cenario.tipoesperado, tipoRecebido)
		}
	}
}
