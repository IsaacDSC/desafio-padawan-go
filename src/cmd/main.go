package main

import (
	"net/http"

	"github.com/IsaacDSC/desafio-padawan-go/src/infra/environments"
	server "github.com/IsaacDSC/desafio-padawan-go/src/infra/server/http"
)

func main() {
	environments.StartEnv(".env")
	http_server := server.HttpServer{}
	server_http := http_server.StartServerHttp()
	http_server.SetMiddleware()
	http_server.SetRouters()
	println("[ * ] Start server http://localhost:3000/")
	if err := http.ListenAndServe(":3000", server_http); err != nil {
		panic(err)
	}
}
