package api

import (
	"fmt"
	"github.com/florestario/adapter/persistence"
	"github.com/florestario/adapter/validator"
	"github.com/florestario/api/controller"
	"github.com/florestario/api/engine"
	"github.com/florestario/core/gateway"
	"github.com/florestario/core/usecase/showroom"
	"goji.io"
	"goji.io/pat"
	"net/http"
)

type server struct {
	mux *goji.Mux
}

type Dependencies struct {
	errorHandler engine.AppErrorHandler
	validatorGateway gateway.ValidatorGateway
	showroomPersistence gateway.ShowroomPersistence
}

func NewApi() *server {
	mux := goji.NewMux()
	server := &server{mux: mux}

	deps := &Dependencies{
		errorHandler:        controller.NewSimpleErrorManager(),
		validatorGateway:    validator.NewOzzoValidator(),
		showroomPersistence: persistence.Instance(),
	}

	server.configureRoutes(deps)
	server.configureFilter(deps)
	server.configureFilterSpecie(deps)

	return server
}

func (server *server) configureFilterSpecie(deps *Dependencies) {
	interactor := showroom.NewFilterSpecieByNameInteractor(deps.showroomPersistence)
	handler := controller.HandlerFilterSpecieByName(interactor)
	server.mux.Handle(pat.Get("/specie/filter"), engine.NewAppContext(deps.errorHandler, handler))
}


func (server *server) configureRoutes(deps *Dependencies) {
	interactor := showroom.NewRegisterGenusInteractor(deps.showroomPersistence, deps.validatorGateway)
	handler := controller.HandlerRegisterGenus(interactor)
	server.mux.Handle(pat.Post("/genus"), engine.NewAppContext(deps.errorHandler, handler))
}

func (server *server) configureFilter(deps *Dependencies) {
	interactor := showroom.NewFilterGenusByNameInteractor(deps.showroomPersistence)
	handler := controller.HandlerFilterGenusByName(interactor)
	server.mux.Handle(pat.Get("/genus/filter"), engine.NewAppContext(deps.errorHandler, handler))
}

func (server *server) Start() {
	fmt.Printf("Server running @ 8080")
	http.ListenAndServe(":8080", server.mux)
}
