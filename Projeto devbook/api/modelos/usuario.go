package modelos

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoem,omitempty"`
}

func (u Usuario) ValidarEntidade(Validasenha bool) error {
	if erro := u.validar(Validasenha); erro != nil {
		return erro
	}

	u.formatar()
	return nil
}

func (u Usuario) validar(Validasenha bool) error {
	var erros []string
	if u.Nome == "" {
		erros = append(erros, "Nome obrigatório e não pode estar em branco")
	}
	if u.Nick == "" {
		erros = append(erros, "Nick obrigatório e não pode estar em branco")
	}
	if u.Email == "" {
		erros = append(erros, "Email obrigatório e não pode estar em branco")
	}

	if checkmail.ValidateFormat(u.Email) != nil {
		erros = append(erros, "E-mail inválido")
	}

	if Validasenha {
		if u.Senha == "" {
			erros = append(erros, "Senha obrigatória e não pode estar em branco")
		}

		if len(u.Senha) < 8 {
			erros = append(erros, "Senha deve ter 8 ou mais caracteres")
		}

		if len(erros) == 0 {
			senhaHash, errosenha := seguranca.Hash(u.Senha)
			if errosenha != nil {
				erros = append(erros, "Não foi possível gerar o hash da senha: "+errosenha.Error())
			} else {
				u.Senha = string(senhaHash)
			}
		}
	}
	if len(erros) != 0 {
		lista := ""
		for _, erro := range erros {
			lista += erro + ","

		}
		lista += "Verificar os campos e corrigir os erros"
		return errors.New(lista)
	}
	return nil
}

func (u *Usuario) formatar() {
	u.Nome = strings.TrimSpace(u.Nome)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
}
