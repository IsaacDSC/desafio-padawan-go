package server

import "github.com/IsaacDSC/desafio-padawan-go/src/infra/server/http/controllers"

func (this_httpServer *HttpServer) SetRouters() {
	this_httpServer.server.Get("/", controllers.HealthController)
	this_httpServer.server.Get("/exchange/{amount}/{from}/{to}/{rate}", controllers.Get_ExchangeRateController)
}
