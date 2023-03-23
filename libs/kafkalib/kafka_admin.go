package kafkalib

import (
	"app/utils"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/spf13/viper"
)

// cgo: C compiler "gcc" not found: exec: "gcc": executable file not found in %PATH% 解决
// https://sourceforge.net/projects/mingw-w64/files/mingw-w64/mingw-w64-release/

var adminClient *kafka.AdminClient

//配置参数：https://github.com/confluentinc/librdkafka/blob/master/CONFIGURATION.md
func InitKafkaClient() {
	host := viper.GetString("kafka.host")
	port := viper.GetString("kafka.port")
	fmt.Println(port)
	conf := &kafka.ConfigMap{
		"bootstrap.servers": host + ":" + port,
		"group.id":          1,
		"auto.offset.reset": "earliest",
	}
	var err error
	adminClient, err = kafka.NewAdminClient(conf)
	if err != nil {
		panic(utils.StringToInterface(err.Error()))
	}
}

func GetClient() *kafka.AdminClient {
	return adminClient
}
