package main

import (
	"app/routers"
	"app/servers"
	"app/utils"
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	loadConfig()
	router := routers.New().Init()
	servers.New(router).Start()

	fmt.Println("server started")
}

// 加载配置文件信息到viper
func loadConfig() {
	envList := []string{"dev", "test", "prod"}
	//env := flag.String("env", "dev", "input run env[dev|test|prod]:")
	defaultEnv := "dev"
	env := &defaultEnv
	flag.Parse()
	if f := utils.InSlice(envList, *env); false == f {
		panic(utils.StringToInterface("env input error"))
	}
	configName := "app.dev"
	switch *env {
	case "dev":
		configName = "app.dev"
	case "test":
		configName = "app.test"
	case "prod":
		configName = "app"
	}
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	if err := viper.ReadInConfig(); err != nil {
		panic(utils.StringToInterface(err.Error()))
	}
	fmt.Println("系统配置如下:")
	fmt.Println("env:", *env)
	fmt.Println("app:", viper.Get("app"))
	fmt.Println("mysql:", viper.Get("mysql"))
	fmt.Println("redis:", viper.Get("redis"))
	fmt.Println("mongo:", viper.Get("mongo"))
	fmt.Println("rabbitmq:", viper.Get("rabbitmq"))
	fmt.Println("elasticsearch:", viper.Get("elasticsearch"))
	fmt.Println("clickhouse:", viper.Get("clickhouse"))
	fmt.Println("kafka:", viper.Get("kafka"))
}
