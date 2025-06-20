package routers

import (
	"modules/src/controllers"
	"net/http"
)

var notes = []Router{
	{
		URI:      "/notes",
		Method:   http.MethodPost,
		Function: controllers.CreateNote,
		Authentication:     true,
	},
	{
		URI:      "/notes",
		Method:   http.MethodGet,
		Function: controllers.GetAllNotes,
		Authentication:     true,
	},
	{
		URI:      "/notes/{noteId}",
		Method:   http.MethodPost,
		Function: controllers.GetByIDNote,
		Authentication:     true,
	},
{
		URI:      "/notes/{noteId}",
		Method:   http.MethodPut,
		Function: controllers.UpdateNote,
		Authentication:     true,
	},
	{
		URI:      "/notes/{noteId}",
		Method:   http.MethodDelete,
		Function: controllers.DeleteNote,
		Authentication:     true,
	},
}