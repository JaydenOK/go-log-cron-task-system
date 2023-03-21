package servers

import "net/http"

type Server struct {
	httpServer http.Server
	host       string
	port       string
}

func New() {

}

func (server *Server) start() {
	//server.httpServer.Handler
	server.httpServer.ListenAndServe(server.host,server.port)
}
