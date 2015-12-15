package main

import "net/http"

type Route struct{
	Name		string
	Method		string
	Pattern		string
	HandlerFunc	http.HandlerFunc
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
		"ApiItemIndex",
		"GET",
		"/api/items",
		ApiItemIndex,
	},
	Route{
		"ApiItem",
		"GET",
		"/api/item/{itemId}",
		ApiItem,
	},
	Route{
		"ApiItemCreate",
		"POST",
		"/api/item",
		ApiItemCreate,
	},
}