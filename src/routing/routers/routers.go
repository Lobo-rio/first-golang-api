package routers

import (
	"modules/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	URI               string
	Method            string
	Function          func(http.ResponseWriter, *http.Request)
	Authentication    bool
}

// Configure function that puts all routes inside the main call
func Configure(r *mux.Router) *mux.Router {
	routers := users
	routers = append(routers, login)

	for _, router := range routers {

		if router.Authentication {
			r.HandleFunc(router.URI,
				middlewares.Logger(middlewares.Authenticate(router.Function)),
			).Methods(router.Method)
		} else {
			r.HandleFunc(router.URI, middlewares.Logger(router.Function)).Methods(router.Method)
		}

	}

	return r
}