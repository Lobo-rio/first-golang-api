package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	URI               string
	Method            string
	Function          func(http.ResponseWriter, *http.Request)
	Authentication    bool
}

// Configurar coloca todas as rotas dentro do router
func Configure(r *mux.Router) *mux.Router {
	routers := users

	for _, router := range routers {

		if router.Authentication {
			r.HandleFunc(router.URI,
				middlewares.Logger(middlewares.Autenticar(router.Function)),
			).Methods(router.Method)
		} else {
			r.HandleFunc(router.URI, middlewares.Logger(router.Function)).Methods(router.Method)
		}

	}

	return r
}