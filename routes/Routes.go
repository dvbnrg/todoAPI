package routes

import (
	"net/http"

	handle "github.com/dvbnrg/todoAPI/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handle.Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		handle.TodoIndex,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		handle.TodoCreate,
	},
	Route{
		"TodoGet",
		"GET",
		"/todos/{todoId}",
		handle.TodoGet,
	},
}
