package main

import (
	"net/http"

	server "github.com/IsaacDSC/desafio-padawan-go/src/infra/server/http"
)

func main() {
	http_server := server.HttpServer{}
	server_http := http_server.StartServerHttp()
	http_server.SetMiddleware()
	http_server.SetRouters()
	println("[ * ] Start server http://localhost:3000/")
	if err := http.ListenAndServe(":3000", server_http); err != nil {
		panic(err)
	}
}
