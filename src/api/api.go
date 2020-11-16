package api

import (
	"goji.io"
	"goji.io/pat"
	"net/http"
)

type server struct {
	mux *goji.Mux
}

func NewApi() *server {
	mux := goji.NewMux()
	server := &server{mux: mux}
	server.configureRoutes()
	return server
}

func (server *server) configureRoutes() {
	server.mux.HandleFunc(pat.Post("/register"), Create)
	server.mux.HandleFunc(pat.Post("/genus"), CreateGenus)
}

func (server *server) Start() {
	http.ListenAndServe(":8080", server.mux)
}
