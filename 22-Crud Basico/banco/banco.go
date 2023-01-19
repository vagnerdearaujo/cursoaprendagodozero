package banco

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //Driver de conexão que deve ser posto explicitamente
)

// ConectarDB Realiza a conexão com o banco de dados, devolve a conexão e erro se houver.
func ConectarDB(dbDriver, dbUser, dbPasswd, dbName, dbServer, connectionParameters string) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s%s", dbUser, dbPasswd, dbServer, dbName, connectionParameters)
	db, erro := sql.Open(dbDriver, connectionString)

	if erro != nil {
		return nil, erro
	}

	//Para verificar se conexão ocorreu com sucesso, existe o método ping

	if erro := db.Ping(); erro != nil {
		return nil, erro
	}
	return db, nil
}
