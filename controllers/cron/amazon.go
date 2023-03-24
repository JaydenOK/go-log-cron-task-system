package cron

import (
	"app/contexts"
	"app/services"
	"app/utils"
	"net/http"
)

type Amazon struct {
	service services.AmazonService
}

//定时任务触发，并发刷新亚马逊token
func (amazon *Amazon) RefreshToken(w http.ResponseWriter, r *http.Request) {
	if data, e := amazon.service.RefreshToken(w, r); e != nil {
		utils.FailResponse(w, e.Error())
	} else {
		utils.SuccessResponse(w, data)
	}
}

//批量拉单
func (amazon *Amazon) PullOrder(w http.ResponseWriter, r *http.Request) {
	amazon.service.PullOrder(contexts.New(w, r))
}

//亚马逊报告
func (amazon *Amazon) GetReportDocument(w http.ResponseWriter, r *http.Request) {
	amazon.service.GetReportDocument(contexts.New(w, r))
}