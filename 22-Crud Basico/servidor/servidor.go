package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Mapeia a tabela usuarios do banco para o struct usuario
type usuarios struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// CriarUsuario Inclui um novo usuário no banco.
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	/*
		Os dados tratados neste método são oriundos diretamente do elemento body
		da página html.
		Os dados estão em formato json e o método ReadAll irá extrair estes dados
	*/
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Dados incorretos enviados para persistência."))

		//O return em um método sem retorno declarado, força seu encerramento
		return
	}

	var usuario usuarios

	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter os dados para persistência"))
		return
	}

	dbUser := "golang_devbook"
	dbPasswd := "devbook_golang"
	dbDriver := "mysql"
	dbName := "devbook"
	dbServer := "172.18.0.2:3306"
	connectionParameters := "?charset=utf8&parseTime=True&loc=Local"
	db, erro := banco.ConectarDB(dbDriver, dbUser, dbPasswd, dbName, dbServer, connectionParameters)

	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	fmt.Println("Conexão realizada com sucesso")

	defer db.Close()

	/*
		Para realizar a inclusão o mais óbvio seria o comando:
		insert into usuarios (nome, email) values ("nome","email")

		o problema desta abordagem é que ela está sujeita ao ataque sql-injection.

		para mitigar este problema existe a preparação do comando em sql.
	*/

	sqlInsert, erro := db.Prepare("insert into usuarios (nome, email) values (?,?)")
	if erro != nil {
		fmt.Println(erro)
		//w.Write([]byte("A preparação do comando de insert falhou"))
		return
	}
	defer sqlInsert.Close()
	//fmt.Println("Comando preparado")

	resultInsert, erro := sqlInsert.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao tentar incluir o usuário."))
		return
	}
	//fmt.Println("Inclusão realizada com sucesso")

	idUsuario, erro := resultInsert.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o id do usuário inserido."))
		return
	}

	fmt.Println(idUsuario, usuario)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("<br>ID:\t %d</br>", idUsuario)))
	w.Write([]byte(fmt.Sprintf("<br>Nome:\t %v</br>", usuario.Nome)))
	w.Write([]byte(fmt.Sprintf("<br>E-Mail:\t %v</br>", usuario.Email)))
}
