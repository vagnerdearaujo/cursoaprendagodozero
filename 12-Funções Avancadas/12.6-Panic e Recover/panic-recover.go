package main

/*
	O panic gera uma interrupção abrupta no que estiver sendo executado.
	Nas linguagens que tem tratamento de erro centralizado, seria o equivalente a executar raise exception
	ou halt.

*/
func main() {
	Println("=========================== Panic & Recover ===========================")

}


func alunoEstaAprovado(n1,n2 float64) bool {
	media := (n1+n2)/2

	if (media > 6) {
		return true
	} else if media < 6 {
		return false
	}

	//Se a média for = 6, gera o panic
	panic ("A média é exatamente 6")

		

}