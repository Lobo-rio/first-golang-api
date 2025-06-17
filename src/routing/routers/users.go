package routers

import "net/http"

var users = []Router{
	{
		URI:                "/users",
		Method:             http.MethodPost,
		Function:             controllers.Create,
		Authentication: false,
	},
	{
		URI:                "/users",
		Method:             http.MethodGet,
		Function:             controllers.getAll,
		Authentication: true,
	},
	{
		URI:                "/users/{userId}",
		Method:             http.MethodGet,
		Function:             controllers.getById,
		Authentication: true,
	},
	{
		URI:                "/users/{userId}",
		Method:             http.MethodPut,
		Function:             controllers.Update,
		Authentication: true,
	},
	{
		URI:                "/users/{usuarioId}",
		Method:             http.MethodDelete,
		Function:             controllers.Delete,
		Authentication: true,
	},
}