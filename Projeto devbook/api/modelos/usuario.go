package modelos

import "time"

type Usuario struct {
	ID       uint      `json:"id,omitempty"`
	Nome     string    `json:"nome, omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty`
	Senha    string    `json:"senha,omitempty`
	CriadoEm time.Time `json:"criadoEm,omitempty`
}
