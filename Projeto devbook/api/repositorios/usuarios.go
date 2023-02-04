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

func (repositorio usuario) ObterUsuario(id uint64) (modelos.Usuario, error) {
	var usuario modelos.Usuario
	resgistro, erro := repositorio.db.Query("select id, nome, nick, email, CriadoEm from usuarios where id = ?", id)
	if erro != nil {
		//Para este caso, não é possível retornar um nil, portanto uma estrutura de usuário precisa ser retornada
		//A variável  usuario é retornada por estar vazia até este ponto
		return modelos.Usuario{}, erro
	}
	defer resgistro.Close()

	if resgistro.Next() {
		erro := resgistro.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm)
		if erro != nil {
			return modelos.Usuario{}, erro
		}
	}
	return usuario, nil
}

func (repositorio usuario) AtualizaDadosUsuario(usuario modelos.Usuario) error {
	statment, erro := repositorio.db.Prepare(`update usuarios set nome = ?,nick = ?,email = ? where id=?`)

	if erro != nil {
		return erro
	}
	defer statment.Close()
	_, erro = statment.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.ID)
	if erro != nil {
		return erro
	}

	return nil
}

func (repositorio usuario) ApagarUsuario(id uint64) error {
	statment, erro := repositorio.db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		return erro
	}
	defer statment.Close()
	_, erro = statment.Exec(id)
	if erro != nil {
		return erro
	}
	return nil
}

func (repositorio usuario) ObterUsuarioPorEmail(email string) (modelos.Usuario, error) {
	var usuario modelos.Usuario
	registro, erro := repositorio.db.Query("select id,senha from usuarios where email = ?", email)
	if erro != nil {
		return usuario, erro
	}
	defer registro.Close()
	if registro.Next() {
		if erro := registro.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return usuario, erro
		}
	} else {
		return usuario, nil //Garante que o zero seja enviado para ser tratado no controller
	}

	return usuario, nil
}
