package loglib

import (
	"app/utils"
	log "github.com/sirupsen/logrus"
	"os"
)

//变量定义
var logger *log.Logger

func Init() *log.Logger {
	logger = log.New()
	logger.SetLevel(log.DebugLevel)
	//显示样式
	logger.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: utils.TimeFormat,
	})
	logger.SetOutput(os.Stdout)
	return logger
}

func GetLogger() *log.Logger {
	return logger
}

func Info(msg string, fields ...log.Fields) {
	fieldsList := log.Fields{}
	for key, value := range fields {
		fieldsList[string(key)] = value
	}
	GetLogger().WithFields(fieldsList).Info(msg)
}
