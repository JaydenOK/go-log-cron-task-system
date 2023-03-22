package servers

import (
	"app/routers"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
)

type Server struct {
	router   *routers.Router
	httpPort string
}

func New(router *routers.Router) *Server {
	var server = &Server{
		router:   router,
		httpPort: viper.GetString("app.httpPort"),
	}
	return server
}

func (server *Server) Start() {
	mux := http.NewServeMux()
	for name, f := range server.router.GetRegisterRouter() {
		mux.HandleFunc(name, f)
	}
	serve := http.ListenAndServe(":"+server.httpPort, mux)
	if serve != nil {
		fmt.Println("启动失败setup fail:", serve)
	} else {
		fmt.Println("success")
	}
}
