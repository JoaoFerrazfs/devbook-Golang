package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasPublicacoes = []Rota{
	{
		URI:                  "/publicacoes",
		Metodo:               http.MethodPost,
		Funcao:               controllers.CriarPublicaco,
		RequerAutentificacao: true,
	},
	{
		URI:                  "/publicacoes",
		Metodo:               http.MethodGet,
		Funcao:               controllers.BuscarPublicacoes,
		RequerAutentificacao: true,
	},
	{
		URI:                  "/publicacoes/{publicaoId}",
		Metodo:               http.MethodGet,
		Funcao:               controllers.BuscarPublicacao,
		RequerAutentificacao: true,
	},
	{
		URI:                  "/publicacoes/{publicaoId}",
		Metodo:               http.MethodPut,
		Funcao:               controllers.AtualizarPublicaco,
		RequerAutentificacao: true,
	},
	{
		URI:                  "/publicacoes/{publicaoId}",
		Metodo:               http.MethodDelete,
		Funcao:               controllers.DeletarPublicaco,
		RequerAutentificacao: true,
	},
	{
		URI:                  "/usuarios/{usuarioId}/publicacoes",
		Metodo:               http.MethodGet,
		Funcao:               controllers.BuscarPublicacoesPorUsuario,
		RequerAutentificacao: true,
	},
	{
		URI:                  "/publicacoes/{publicaoId}/curtir",
		Metodo:               http.MethodPost,
		Funcao:               controllers.CurtirPublicacao,
		RequerAutentificacao: true,
	},
	{
		URI:                  "/publicacoes/{publicaoId}/descurtir",
		Metodo:               http.MethodPost,
		Funcao:               controllers.DescurtirPublicacao,
		RequerAutentificacao: true,
	},
}
