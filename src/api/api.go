package api

import (
	"fmt"
	"github.com/florestario/api/controller"
	"github.com/florestario/api/engine"
	"github.com/florestario/core/mock"
	"github.com/florestario/core/usecase/showroom"
	"github.com/florestario/persistence"
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
	pg := persistence.NewAquaticPostgres()
	interactor := showroom.NewRegisterGenusInteractor(pg, mock.NewValidatorMock())
	genusCtr := controller.NewGenusController(interactor)
	server.mux.Handle(pat.Post("/genus"), engine.NewAppContext(errorManager, genusCtr.HandleCreateGenus))
}

func (server *server) Start() {
	fmt.Printf("Server running @ 8080")
	http.ListenAndServe(":8080", server.mux)
}
