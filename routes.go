package main

import "net/http"

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
		Index,
	},
	Route{
		"LocationCreate",
		"POST",
		"/location",
		LocationCreate,
	},
	Route{
		"LocationShow",
		"GET",
		"/location/{location_id}",
		LocationShow,
	},
	Route{
		"LocationUpdate",
		"PUT",
		"/location/{location_id}",
		LocationUpdate,
	},
	Route{
		"LocationRemove",
		"DELETE",
		"/location/{location_id}",
		LocationRemove,
	},

}
