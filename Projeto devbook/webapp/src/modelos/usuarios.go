package modelos

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requisicoes"
)

// Usuario: Estrutura de dados para listagem de usuários e perfil do usuário logado
type Usuario struct {
	ID          uint64       `json:"id"`
	Nome        string       `json:"nome"`
	Email       string       `json:"email"`
	Nick        string       `json:"nick"`
	CriadoEm    time.Time    `json:"criadoEm"`
	Seguidores  []Usuario    `json:"seguidores"`
	Seguindo    []Usuario    `json:"seguindo"`
	Publicacoes []Publicacao `json:"publicacoes"`
}

// CarregarPerfilUsuario: Recebe o Id do usuário e a requisição e devolve o perfil completo do usuário
func CarregarPerfilUsuario(UsuarioId uint64, r *http.Request) (Usuario, error) {
	//Configurar os canais de comunicação para cada um dos tipos de dados do usuário
	canalUsuario := make(chan Usuario)
	canalSeguidores := make(chan []Usuario)
	canalSeguindo := make(chan []Usuario)
	canalPublicacoes := make(chan []Publicacao)

	go BuscarDadosUsuario(canalUsuario, UsuarioId, r)
	go BuscarSeguidores(canalSeguidores, UsuarioId, r)
	go BuscarSeguindo(canalSeguindo, UsuarioId, r)
	go BuscarPublicacoes(canalPublicacoes, UsuarioId, r)

	var (
		usuario     Usuario
		seguidores  []Usuario
		seguindo    []Usuario
		publicacoes []Publicacao
	)

	for i := 0; i < 4; i++ {
		select {
		case usuarioCarregado := <-canalUsuario:
			if usuarioCarregado.ID == 0 {
				return Usuario{}, errors.New("Erro ao buscar o usuário")
			}
			usuario = usuarioCarregado
		case seguidoresCarregados := <-canalSeguidores:
			if seguidoresCarregados == nil {
				return Usuario{}, errors.New("Erro ao carregar lista de seguidores")
			}
			seguidores = seguidoresCarregados
		case seguindoCarregados := <-canalSeguindo:
			if seguindoCarregados == nil {
				return Usuario{}, errors.New("Erro ao carregar lista de quem está sendo seguido")
			}
			seguindo = seguindoCarregados
		case publicacoesCarregadas := <-canalPublicacoes:
			if publicacoesCarregadas == nil {
				return Usuario{}, errors.New("Erro ao carregar lista de publicacoes")
			}
			publicacoes = publicacoesCarregadas
		}
	}

	usuario.Seguidores = seguidores
	usuario.Seguindo = seguindo
	usuario.Publicacoes = publicacoes

	return usuario, nil
}

// BuscarDadosUsuario: Recupera os dados básicos do usuário e os envia através do canal
func BuscarDadosUsuario(canal chan<- Usuario, usuarioId uint64, r *http.Request) {
	urlAPI := config.APIAddress(fmt.Sprintf("usuarios/%d", usuarioId))
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlAPI, nil)
	if erro != nil {
		//Se houver erro envia um usuário vazio para o canal e encerra
		canal <- Usuario{}
		return
	}

	defer response.Body.Close()

	var usuario Usuario
	if erro = json.NewDecoder(response.Body).Decode(&usuario); erro != nil {
		canal <- Usuario{}
		return
	}

	canal <- usuario
}

// BuscarSeguidores: Recupera a lista seguidores do usuárioos e a envia através do canal
func BuscarSeguidores(canal chan<- []Usuario, usuarioId uint64, r *http.Request) {
	urlAPI := config.APIAddress(fmt.Sprintf("usuarios/%d/seguidores", usuarioId))
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlAPI, nil)
	if erro != nil {
		//Se houver erro envia um usuário vazio para o canal e encerra
		canal <- nil
		return
	}

	defer response.Body.Close()

	var seguidores []Usuario
	if erro = json.NewDecoder(response.Body).Decode(&seguidores); erro != nil {
		canal <- nil
		return
	}

	if seguidores == nil {
		canal <- make([]Usuario, 0)
		return
	}
	canal <- seguidores
}

// BuscarSeguindo: Recupera a lista de quem o usuário segue e a envia através do canal
func BuscarSeguindo(canal chan<- []Usuario, usuarioId uint64, r *http.Request) {
	urlAPI := config.APIAddress(fmt.Sprintf("usuarios/%d/segue", usuarioId))
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlAPI, nil)
	if erro != nil {
		canal <- nil
		return
	}

	defer response.Body.Close()

	var seguindo []Usuario
	if erro = json.NewDecoder(response.Body).Decode(&seguindo); erro != nil {
		canal <- nil
		return
	}
	if seguindo == nil {
		canal <- make([]Usuario, 0)
		return
	}
	canal <- seguindo
}

// BuscarPublicacoes: Recupera a lista de publicações relacionadas ao usuário e a envia através do canal
func BuscarPublicacoes(canal chan<- []Publicacao, usuarioId uint64, r *http.Request) {

	urlAPI := config.APIAddress(fmt.Sprintf("/publicacoes/%d/usuarios", usuarioId))
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, urlAPI, nil)
	if erro != nil {
		canal <- nil
		return
	}

	defer response.Body.Close()

	var publicacoes []Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
		canal <- nil
		return
	}

	if publicacoes == nil {
		canal <- make([]Publicacao, 0)
		return
	}
	canal <- publicacoes
}
