package daemon

import (
	"github.com/gorilla/mux"
)

func initRoutes() *mux.Router {
	m := mux.NewRouter()
	m.Handle("/pong", PongHandler()).Methods("GET")
	return m
}
