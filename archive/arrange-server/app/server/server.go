package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/hueyjj/arrange-server/controller"
)

type Server struct {
	Router *mux.Router
}

func NewServer() *Server {
	server := &Server{
		Router: mux.NewRouter(),
	}

	controller.Init(server.Router)

	return server
}

func (s *Server) Shutdown() {
}

func (s *Server) Start() error {
	log.Info("Starting server...")

	srv := &http.Server{
		Handler:      s.Router,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Error(err)
		}
	}()

	return nil
}
