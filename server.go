package matchfinder

import (
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(port string, router *http.ServeMux) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + port,
			Handler:        router,
			MaxHeaderBytes: 1 << 20,
			ReadTimeout:    time.Second * 10,
			WriteTimeout:   time.Second * 10,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}
