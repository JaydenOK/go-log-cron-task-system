package routers

import (
	"app/controllers/api"
	"app/controllers/cron"
	"net/http"
)

type Router struct {
	registerRouter map[string]func(http.ResponseWriter, *http.Request)
}

func New() *Router {
	var router = &Router{
		registerRouter: make(map[string]func(w http.ResponseWriter, r *http.Request)),
	}
	return router
}

//路由配置
func (r *Router) urlMapping() {
	{
		var cmd api.Cmd
		r.Add("/api/cmd/start", cmd.Start)
		r.Add("/api/cmd/stop", cmd.Stop)
		r.Add("/api/cmd/status", cmd.Status)
	}
	{
		var kafkaManage api.KafkaManage
		r.Add("/api/es/topic_create", kafkaManage.TopicCreate)
		r.Add("/api/es/topic_list", kafkaManage.TopicList)
	}
	{
		var es api.Es
		r.Add("/api/es/add_index", es.AddIndex)
		r.Add("/api/es/search", es.Search)
	}
	{
		//定时任务
		var amazon cron.Amazon
		r.Add("/cron/amazon/refresh_token", amazon.RefreshToken)
		r.Add("/cron/amazon/pullOrder", amazon.PullOrder)
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
	return r
}
