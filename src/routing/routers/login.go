package routers

import (
	"modules/src/controllers"
	"net/http"
)

var login = Router{
	URI:            "/login",
	Method:         http.MethodPost,
	Function:       controllers.Login,
	Authentication: false,
}