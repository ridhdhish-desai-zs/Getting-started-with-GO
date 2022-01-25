package routes

import (
	"net/http"
	"webserver/controller"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var UserRoutes = Routes{
	Route{
		"Users",
		"GET",
		"/api/users",
		controller.UserRequestHandler,
	},
	Route{
		"UserIndex",
		"GET",
		"/api/users/{id}",
		controller.FindUserByIdRequestHandler,
	},
	Route{
		"UserCreate",
		"POST",
		"/api/users",
		controller.CreateUserRequestHandler,
	},
}
