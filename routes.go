package main

import (
  "net/http"
)

type RouteHandler func(http.ResponseWriter, *http.Request)

type Route struct {
  Method      string
	Path        string
	HandlerFunc RouteHandler
}

type Routes []Route

func AllRoutes() Routes {
	routes := Routes{
    Route{ "POST", "/give", Give },
    Route{ "POST", "/get-score", GetScore },
    Route{ "POST", "/all", GetAllScores },
    Route{ "POST", "/rem", GetAllRemaining },
	}
	return routes
}
