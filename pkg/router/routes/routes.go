package routes

import (
	"devbook-api/pkg/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Uri            string
	Method         string
	Function       func(http.ResponseWriter, *http.Request)
	Authentication bool
}

func Configurate(r *mux.Router) *mux.Router {
	allRoutes := userRoutes
	allRoutes = append(allRoutes, routeLogin)

	for _, route := range allRoutes {
		if route.Authentication {
			r.HandleFunc(route.Uri, middleware.Authenticate(route.Function)).Methods(route.Method)
		} else {
			r.HandleFunc(route.Uri, route.Function).Methods(route.Method)
		}
	}

	return r
}
