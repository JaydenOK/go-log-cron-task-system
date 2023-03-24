package api

import (
	"net/http"
)

//kafka-sdk文档：https://docs.confluent.io/platform/current/clients/confluent-kafka-go/index.html

type KafkaManage struct {
}

func (kafkaManage *KafkaManage) TopicCreate(w http.ResponseWriter, r *http.Request) {

}

func (kafkaManage *KafkaManage) TopicList(w http.ResponseWriter, r *http.Request) {
	//adminClient := kafkalib.GetClient()
	//metadata, err := adminClient.GetMetadata(nil, true, 100)
	//if err != nil {
	//	utils.FailResponse(w, err)
	//	return
	//}
	//for key, value := range metadata.Topics {
	//	fmt.Println("key:"+key, "value:", value)
	//}
	//utils.SuccessResponse(w, metadata.Topics)
}
