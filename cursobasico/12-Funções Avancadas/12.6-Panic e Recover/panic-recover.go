package main

/*
O panic gera uma interrupção abrupta no que estiver sendo executado.
Nas linguagens que tem tratamento de erro centralizado, seria o equivalente a executar halt.
Após a execução do panic o programa morre!!!!.
Vai tentar executar os métodos apontados por defer e encerra

Não é a melhor forma de tratar uma situação de erro, pois interrompe a execução do programa
*/
func main() {
	println("=========================== Panic & Recover ===========================")
	println("Resultado de Aluno aprovado = ",alunoEstaAprovado(6, 6))
	println("Pós execução")
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperaExecucao()
	defer println("Print via defer!!!!") //Esta linha irá aparecer, mesmo que haja o panic
	println("Calculando a média.")
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	//Se a média for = 6, gera o panic
	panic("A média é exatamente 6")
}

/*
O panic executa os métodos adiados com defer antes de encerrar o programa, portanto, neste ponto
é possível tentar uma recuperação com o comando recover.

O recover não irá gerar erro caso não haja a necessidade de recuperação.
*/
func recuperaExecucao() {
	if r := recover(); r != nil {
		println("Execução recuperada com sucesso !!!!!!")
	}
}
