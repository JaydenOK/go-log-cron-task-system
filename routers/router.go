package routers

import (
	"app/controllers"
	"net/http"
)

type Router struct {
	registerRouter map[string]func(http.ResponseWriter, *http.Request)
}

var router = &Router{
	registerRouter: make(map[string]func(w http.ResponseWriter, r *http.Request)),
}

func New() *Router {
	return router
}

//路由配置
func (r *Router) urlMapping() {
	{
		var cmd controllers.Cmd
		r.Add("/cmd/start", cmd.Start)
		r.Add("/cmd/stop", cmd.Stop)
		r.Add("/cmd/status", cmd.Status)
	}
	{
		var es controllers.KafkaManage
		r.Add("/es/create_topic", es.AddIndex)
		r.Add("/es/search", es.Search)
	}
	{
		var es controllers.Es
		r.Add("/es/add_index", es.AddIndex)
		r.Add("/es/search", es.Search)
	}
}

func (r *Router) Add(name string, f func(http.ResponseWriter, *http.Request)) {
	_, ok := r.registerRouter[name]
	if !ok {
		r.registerRouter[name] = f
	}
}

func (r *Router) Remove(name string) {
	_, ok := r.registerRouter[name]
	if ok {
		delete(r.registerRouter, name)
	}
}

func (r *Router) GetRegisterRouter() map[string]func(http.ResponseWriter, *http.Request) {
	return r.registerRouter
}

//初始化
func (r *Router) Init() *Router {
	r.urlMapping()
	return router
}
