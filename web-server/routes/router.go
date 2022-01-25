package routes

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	routes := UserRoutes

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		// var handler http.Handler

		// handler = route.HandlerFunc
		// handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)

	}

	return router
}
