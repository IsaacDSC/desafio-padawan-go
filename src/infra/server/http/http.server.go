package server

import (
	"github.com/go-chi/chi"
)

type HttpServer struct {
	server *chi.Mux
}

func (this_httpServer *HttpServer) StartServerHttp() *chi.Mux {
	this_httpServer.server = chi.NewRouter()
	return this_httpServer.server
}
