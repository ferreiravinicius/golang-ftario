package api

import (
	"fmt"
	"github.com/florestario/api/controller"
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
	server.mux.HandleFunc(pat.Post("/genus"), controller.CreateGenus(nil).ServeHTTP)
}

func (server *server) Start() {
	fmt.Printf("Server running @ 8080")
	http.ListenAndServe(":8080", server.mux)
}
