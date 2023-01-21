package main

import "fmt"

/*
	O Map é uma estrutura do tipo chave/valor.
	Esta estrutura pode ser entendida como dictionary do Delphi.

	A estrutura exige que se especifique o tipo da chave e o tipo do valor e uma especificado isto, as chaves e valores
	obrigatoriamente devem respeitar os tipos especificados.
*/

func main() {
	fmt.Println("=========================== Arrays ===========================")
	usuario := map[string]string{
		"nome":      "Vagner",
		"sobrenome": "de Araujo",
		"fone":      "(11) 992.314.925",
	}

	patrimonios := map[string]int{
		"notebook lenovo":  1,
		"monitor philips":  2,
		"monitor proview":  3,
		"teclado logitech": 4,
	}

	fmt.Println(usuario)
	fmt.Println(patrimonios)

	/*
		Para acessar os elementos, deve-se agir como se o map fosse um array
		Se uma chave for especificada e não existir não há retorno
		Se um valor for especificado para uma chave que originalmente não foi declarado, passa a existir.
	*/

	fmt.Printf("Dados do Usuário:\nNome:%v %v\nTelefone:%v\n", usuario["nome"], usuario["sobrenome"], usuario["fone"])
	fmt.Printf("Chave que não existe: %v\n", usuario["naoexiste"])

	usuario["naoexiste"] = "Agora existe"
	fmt.Printf("Chave que não existia: %v\n", usuario["naoexiste"])

	/*
		Os maps podem ser aninhados
	*/

	usuario2 := map[string]map[string]string{
		"dados": {
			"nome":      "Vagner",
			"sobrenome": "de Araujo",
			"profissao": "Coach",
		},
		"contato": {
			"endereco": "Rua X, 121",
			"telefone": "(11) 992.314.925",
		},
		"comercial": {
			"empresa": "MPS",
			"cidade":  "São Paulo",
		},
	}

	fmt.Println(usuario2)
	fmt.Println("Telefone: " + usuario2["contato"]["telefone"])

	//Para excluir uma chave basta utilizar o comando delete
	delete(usuario2, "comercial")
	fmt.Println(usuario2)

	//Para incluir uma chave composta é preciso informar a estrutura antes
	usuario2["astral"] = map[string]string{"signo": "Capricórnio"}
	fmt.Println(usuario2)
}
