package routers

import "net/http"

var rotasUsuarios = []Router{
	{
		URI:                "/usuarios",
		Method:             http.MethodPost,
		Function:             controllers.CriarUsuario,
		Authentication: false,
	},
	{
		URI:                "/usuarios",
		Method:             http.MethodGet,
		Function:             controllers.BuscarUsuarios,
		Authentication: true,
	},
	{
		URI:                "/usuarios/{usuarioId}",
		Method:             http.MethodGet,
		Function:             controllers.BuscarUsuario,
		Authentication: true,
	},
	{
		URI:                "/usuarios/{usuarioId}",
		Method:             http.MethodPut,
		Function:             controllers.AtualizarUsuario,
		Authentication: true,
	},
	{
		URI:                "/usuarios/{usuarioId}",
		Method:             http.MethodDelete,
		Function:             controllers.DeletarUsuario,
		Authentication: true,
	},
}