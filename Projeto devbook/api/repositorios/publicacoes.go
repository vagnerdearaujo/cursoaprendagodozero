package repositorios

import (
	"api/modelos"
	"database/sql"
)

type publicacao struct {
	db *sql.DB
}

func NovoRepositorioPublicacao(db *sql.DB) *publicacao {
	return &publicacao{db}
}

func (repositorio publicacao) IncluirrPublicacao(publicacao modelos.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare("insert into publicacoes (titulo,conteudo,autorId) values (?,?,?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	publicacaoId, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(publicacaoId), nil

}
func (repositorio publicacao) AtualizarPublicacao(publicacao modelos.Publicacao) error {
	statement, erro := repositorio.db.Prepare("update publicacoes set titulo = ?,conteudo = ?,where id = publicacaoId")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.ID)
	return erro
}
func (repositorio publicacao) ExcluirPublicacao(publicacaoId uint64) error {
	statement, erro := repositorio.db.Prepare("delete from publicacoes where id = publicacaoId")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(publicacaoId)
	return erro
}

func (repositorio publicacao) ListarPublicacaoId(publicacaoId uint64) (modelos.Publicacao, error) {
	query := `select pub.id,
					 pub.titulo,
					 pub.conteudo,
					 pub.autorId,
					 usr.Nick,
					 pub.curtidas,
					 pub.criadaEm
			   from publicacoes pub
			   inner join usuarios usr on pub.autorId = usr.id
			   where pub.id = ?`
	registro, erro := repositorio.db.Query(query, publicacaoId)
	if erro != nil {
		return modelos.Publicacao{}, erro
	}
	defer registro.Close()
	var publicacao modelos.Publicacao
	if registro.Next() {
		if erro := registro.Scan(&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.AutorNick,
			&publicacao.Curtidas,
			&publicacao.CriadaEm); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

func (repositorio publicacao) ListarPublicacoes(usuarioId uint64) ([]modelos.Publicacao, error) {
	query := `select distinct
					 pub.id,
					 pub.titulo,
					 pub.conteudo,
					 pub.autorId,
					 usr.Nick,
					 pub.curtidas,
					 pub.criadaEm
			   from publicacoes pub
			   inner join usuarios usr on pub.autorId = usr.id
			   inner join seguidores seg on pub.autorId = seg.usuario_id
			   where usr.Id = ? or seg.seguidor_id = ?
			   order by pub.id desc`
	registro, erro := repositorio.db.Query(query, usuarioId, usuarioId)
	if erro != nil {
		return nil, erro
	}
	defer registro.Close()
	var publicacoes []modelos.Publicacao
	for registro.Next() {
		var publicacao modelos.Publicacao
		if erro := registro.Scan(&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.AutorNick,
			&publicacao.Curtidas,
			&publicacao.CriadaEm); erro != nil {
			return nil, erro
		}
		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}
