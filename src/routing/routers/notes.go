package routers

import (
	"modules/src/controllers"
	"net/http"
)

var notes = []Router{
	{
		URI:      "/notes",
		Method:   http.MethodPost,
		Function: controllers.Create,
		Authentication:     true,
	},
	{
		URI:      "/notes",
		Method:   http.MethodGet,
		Function: controllers.GetAll,
		Authentication:     true,
	},
	{
		URI:      "/notes/{noteId}",
		Method:   http.MethodPost,
		Function: controllers.GetById,
		Authentication:     true,
	},
{
		URI:      "/notes/{noteId}",
		Method:   http.MethodPut,
		Function: controllers.Update,
		Authentication:     true,
	},
	{
		URI:      "/notes/{noteId}",
		Method:   http.MethodDelete,
		Function: controllers.Delete,
		Authentication:     true,
	},
}