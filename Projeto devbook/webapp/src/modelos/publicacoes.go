package modelos

import "time"

// Publicacao Modelo para receber os dados de publicação de usuário vinda da API
type Publicacao struct {
	ID        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorID   uint64    `json:"autorid,omitempty"`
	AutorNick string    `json:"autornick,omitempty`
	Curtidas  uint64    `json:"curtidas"`
	CriadaEm  time.Time `json:criadaem,"omitempty"`
}
