package routes

import (
	userHandler "github.com/user-management-system/cmd/handlers/user"
	router "github.com/user-management-system/pkg/router"
)

func GetRoutes() []router.Route {
	type Routes []router.Route

	var routes = Routes{
		router.Route{
			"Index",
			"GET",
			"/",
			userHandler.Index,
		},
		router.Route{
			"UserIndex",
			"GET",
			"/users",
			userHandler.UserIndex,
		},
		router.Route{
			"UserGet",
			"GET",
			"/users/{userId:[0-9]+}",
			userHandler.UserGet,
		},
		router.Route{
			"UserCreate",
			"POST",
			"/users",
			userHandler.UserCreate,
		},
	}

	return routes
}
