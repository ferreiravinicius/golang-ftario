package delivery

import (
	"github.com/florestario/delivery/controller"
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

//func register(w http.ResponseWriter, r *http.Request) error {
//	var genusInput entity.Genus
//	if err := Decode(r, &genusInput); err != nil {
//		return err
//	}
//
//	pg := persistence.NewAquaticPostgres()
//	genusService := service.NewGenusService(pg)
//	output, err := genusService.CreateGenus(&genusInput)
//	if err != nil {
//		return err
//	}
//	AnswerSuccess(w, http.StatusAccepted, output)
//
//	return nil
//}

func (server *server) configureRoutes() {
	server.mux.HandleFunc(pat.Post("/register"), Create)
	server.mux.HandleFunc(pat.Post("/genus"), controller.RegisterGenus.ServeHTTP)
}

func (server *server) Start() {
	http.ListenAndServe(":8080", server.mux)
}
