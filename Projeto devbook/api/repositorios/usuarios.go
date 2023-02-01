package repositorios

import (
	"api/modelos"
	"database/sql"
	"fmt"
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

// Lista todos os usuários que atendam um determinado nick ou name
func (repositorio usuario) ListarUsuarios(nickouname string) ([]modelos.Usuario, error) {
	nickouname = fmt.Sprintf("%%%s%%", nickouname) //%% indica para printf que é para considerar como % literal
	records, erro := repositorio.db.Query("select id, nome, nick, email, CriadoEm from usuarios where lower(nome) like ? or lower(nick) like ?", nickouname, nickouname)
	if erro != nil {
		return nil, erro
	}
	defer records.Close()

	var usuarios []modelos.Usuario
	for records.Next() {
		var usuario modelos.Usuario
		if erro := records.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}
