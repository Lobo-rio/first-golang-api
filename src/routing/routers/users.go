package routers

import (
	"modules/src/controllers"
	"net/http"
)

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
		Function:             controllers.GetAll,
		Authentication: true,
	},
	{
		URI:                "/users/{userId}",
		Method:             http.MethodGet,
		Function:             controllers.GetById,
		Authentication: true,
	},
	{
		URI:                "/users/{userId}",
		Method:             http.MethodPut,
		Function:             controllers.Update,
		Authentication: true,
	},
	{
		URI:                "/users/{userId}",
		Method:             http.MethodDelete,
		Function:             controllers.Delete,
		Authentication: true,
	},
}