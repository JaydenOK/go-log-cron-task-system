package controllers

import (
	"app/messages"
	"app/utils"
	"fmt"
	"net/http"
)

type KafkaManage struct {
}

func (kafkaManage *KafkaManage) Status(w http.ResponseWriter, r *http.Request) {
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
