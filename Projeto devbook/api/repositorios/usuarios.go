package repositorios

import (
	"api/modelos"
	"database/sql"
	"errors"
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

// ObterUsuario Obtém um usuário buscando por Id
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

func (repositorio usuario) SeguirUsuario(SeguidorID, SeguidoID uint64) (modelos.Usuario, error) {
	usuarioSeguido, erro := repositorio.ObterUsuario(SeguidoID)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	if usuarioSeguido.ID == 0 {
		return modelos.Usuario{}, errors.New("Usuaário inexistente")
	}
	statatement, erro := repositorio.db.Prepare("insert into seguidores (usuario_id,seguidor_id) values (?,?)")
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer statatement.Close()
	if _, erro := statatement.Exec(SeguidorID, SeguidoID); erro != nil {
		return modelos.Usuario{}, erro
	}
	return usuarioSeguido, nil

}

func (repositorio usuario) PararSeguir(seguidorID, seguidoID uint64) (bool, error) {
	usuarioSeguido, erro := repositorio.ObterUsuario(seguidoID)
	if erro != nil {
		return false, erro
	}
	if usuarioSeguido.ID == 0 {
		return false, errors.New("Usuaário inexistente")
	}
	statement, erro := repositorio.db.Prepare("delete from seguidores where usuario_id = ? and seguidor_id = ?")
	if erro != nil {
		return false, erro
	}
	defer statement.Close()

	records, erro := statement.Exec(seguidorID, seguidoID)
	if rows, erro := records.RowsAffected(); rows < 1 || erro != nil {
		if rows < 1 {
			erro = errors.New("Você não seguia o usuário: " + usuarioSeguido.Nome)
		}
		return false, erro
	}

	return true, nil
}

func (repositorio usuario) ObterSeguidores(usuarioID uint64) ([]modelos.Usuario, error) {
	strquery := `select usr.id
					, usr.nome
					, usr.nick
					, usr.email
					, usr.criadoem
				from usuarios usr
				inner join seguidores seg on usr.id = seg.seguidor_id
				where seg.usuario_id = ?`

	var seguidores []modelos.Usuario
	registros, erro := repositorio.db.Query(strquery, usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer registros.Close()

	for registros.Next() {
		var usuario modelos.Usuario
		erro := registros.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm)
		if erro != nil {
			return nil, erro
		}
		seguidores = append(seguidores, usuario)
	}

	return seguidores, nil
}

func (repositorio usuario) ObterSeguidos(usuarioID uint64) ([]modelos.Usuario, error) {
	strquery := `select usr.id
					, usr.nome
					, usr.nick
					, usr.email
					, usr.criadoem
				from usuarios usr
				inner join seguidores seg on usr.id = seg.usuario_id
				where seg.seguidor_id = ?`

	var seguidos []modelos.Usuario
	registros, erro := repositorio.db.Query(strquery, usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer registros.Close()

	for registros.Next() {
		var usuario modelos.Usuario
		erro := registros.Scan(&usuario.ID, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm)
		if erro != nil {
			return nil, erro
		}
		seguidos = append(seguidos, usuario)
	}

	return seguidos, nil
}

func (repositorio usuario) AtualizaSenha(usuarioID uint64, SenhaNova string) error {
	statment, erro := repositorio.db.Prepare("update usuarios set senha=? where id=?")
	if erro != nil {
		return erro
	}
	defer statment.Close()
	_, erro = statment.Exec(SenhaNova, usuarioID)
	return erro
}

func (repositorio usuario) ObtemSenhaAtual(usuarioID uint64) (string, error) {
	registro, erro := repositorio.db.Query("select senha from usuarios where id=?", usuarioID)
	if erro != nil {
		return "", erro
	}
	defer registro.Close()
	var senhaAtual string
	if registro.Next() {
		if erro := registro.Scan(&senhaAtual); erro != nil {
			return "", erro
		}
	}

	return senhaAtual, nil
}
