package api

import (
	"fmt"
	"github.com/florestario/api/controller"
	"github.com/florestario/api/engine"
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
	errorManager := controller.NewSimpleErrorManager()
	server.configureRoutes(errorManager)
	return server
}

func (server *server) configureRoutes(errorManager engine.AppErrorHandler) {
	appHandler := controller.CreateGenus(nil)
	server.mux.Handle(pat.Post("/genus"), engine.NewAppController(errorManager, appHandler))
}

func (server *server) Start() {
	fmt.Printf("Server running @ 8080")
	http.ListenAndServe(":8080", server.mux)
}
