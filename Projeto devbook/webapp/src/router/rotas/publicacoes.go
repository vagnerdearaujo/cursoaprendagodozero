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
	{
		//Carrega a página com os dados da publicação para edição
		URI:                "/publicacoes/{publicacaoId}/editar",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDePublicacao,
		RequerAutenticacao: true,
	},
	{
		//Efetiva a edição dos dados de publicação
		URI:                "/publicacoes/{publicacaoId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.EditarPublicacao,
		RequerAutenticacao: true,
	},
	{
		//Excluir a publicação
		URI:                "/publicacoes/{publicacaoId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.ExcluirPublicacao,
		RequerAutenticacao: true,
	},
}
