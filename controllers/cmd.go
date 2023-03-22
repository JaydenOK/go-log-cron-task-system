package controllers

import (
	"app/messages"
	"app/utils"
	"fmt"
	"net/http"
)

type Cmd struct {
}

func (cmd *Cmd) New() {

}

func (cmd *Cmd) Start(w http.ResponseWriter, r *http.Request) {

}

func (cmd *Cmd) Stop(w http.ResponseWriter, r *http.Request) {

}

func (cmd *Cmd) Status(w http.ResponseWriter, r *http.Request) {
	fmt.Println(utils.RequestHeader(r))
	fmt.Println(utils.ParamsGet(r))
	fmt.Println(utils.ParamsPost(r))
	fmt.Println(utils.ParamsJson(r))
	message := messages.SearchMessage{
		Id:         "111",
		Content:    "123ABC呐呐呐",
		CreateTime: "2023-03-22 18:00:00",
	}
	utils.SuccessResponse(w, message)
}
