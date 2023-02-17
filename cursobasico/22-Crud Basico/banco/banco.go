package banco

import (
	"crud/settings"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //Driver de conexão que deve ser posto explicitamente
)

// ConectarDB Realiza a conexão com o banco de dados, devolve a conexão e erro se houver.
func ConectarDB(mySQLSettings settings.DBSettings) (*sql.DB, error) {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s%s", mySQLSettings.DBUser, mySQLSettings.DBPasswd, mySQLSettings.DBServer, mySQLSettings.DBName, mySQLSettings.ConnectionParameters)
	db, erro := sql.Open(mySQLSettings.DBDriver, connectionString)

	if erro != nil {
		return nil, erro
	}

	//Para verificar se conexão ocorreu com sucesso, existe o método ping

	if erro := db.Ping(); erro != nil {
		return nil, erro
	}
	return db, nil
}
