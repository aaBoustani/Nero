package main

import (
  "github.com/gorilla/mux"
)

func setupRouter(routes Routes) *mux.Router {
  router := mux.NewRouter()

  for _, route := range routes {
		handle := Logger(route.HandlerFunc)

		router.HandleFunc(route.Path, handle).Methods(route.Method)
}

  return router
}
