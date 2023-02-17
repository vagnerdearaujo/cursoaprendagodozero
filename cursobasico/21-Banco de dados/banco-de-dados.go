package main

/*
	Usando a importação padrão, teremos um erro em tempo de execução:
 	sql: unknown driver "mysql" (forgotten import?)

	Para evitar este erro, a referência do pacote importada e que está no go.mod deve ser colocada manualmente
	precedida por um underline "_"
*/

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbUser := "golang_devbook"
	dbPasswd := "devbook_golang"
	dbDriver := "mysql"
	dbName := "devbook"
	dbServer := "172.18.0.2:3306"
	connectionParameters := "?charset=utf8&parseTime=True&loc=Local"

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s%s", dbUser, dbPasswd, dbServer, dbName, connectionParameters)
	fmt.Println("ConnectionString:" + connectionString)
	db, erro := sql.Open(dbDriver, connectionString)

	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close() //Garante que a conexão será fechada como última instrução desta função.

	//Para verificar se conexão ocorreu com sucesso, existe o método ping

	if erro := db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conexão realizada com sucesso.")

	tuplas, erro := db.Query("select * from usuarios")
	if erro != nil {
		log.Fatal(erro)
	}
	defer tuplas.Close()

	fmt.Println(tuplas)
}
