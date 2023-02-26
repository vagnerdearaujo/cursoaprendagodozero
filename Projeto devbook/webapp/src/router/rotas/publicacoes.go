package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasPublicacoes = []WebRotas{
	{
		URI:                "/publicacoes",
		Funcao:             controllers.NovaPublicacao,
		Metodo:             http.MethodPost,
		RequerAutenticacao: true,
	},
}

/*
Rotas copiadas da API

	{
		URI:                "/publicacoes/{publicacaoId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes/{publicacaoId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.ExcluirPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes",
		Metodo:             http.MethodGet,
		Funcao:             controllers.ListarPublicacoes,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes/{publicacaoId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPublicacaoId,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes/{usuarioId}/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPublicacaoUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes/{publicacaoId}/curtir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CurtirPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes/{publicacaoId}/descurtir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.DescurtirPublicacao,
		RequerAutenticacao: true,
	},


*/
