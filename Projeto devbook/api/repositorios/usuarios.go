package repositorios

import (
	"api/modelos"
	"database/sql"
)

type usuario struct {
	db *sql.DB
}

func NovoRepositorioUsuario(db *sql.DB) *usuario {
	return &usuario{db}
}

// NovoUsuario retorna o ID do usuário quando sucesso e erro quando fracasso
func (repositorio usuario) NovoUsuario(usuario modelos.Usuario) (uint64, error) {
	statment, erro := repositorio.db.Prepare(`insert into usuarios (nome,nick,email,Senha) values (?,?,?,?)`)

	if erro != nil {
		return 0, erro
	}
	defer statment.Close()
	resultado, erro := statment.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	id, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(id), nil

}
