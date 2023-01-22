package banco

import (
	"api/src/router/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// ConectarBanco abre e retorna a conexão com o banco ou erro caso ocorra
func ConectarBanco() (*sql.DB, error) {
	db, erro := sql.Open(config.DB_driverbanco, config.StringConexaoBanco)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
