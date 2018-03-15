package general_log

import (
	"fmt"
	"log"

	"github.com/itziklavon/kit2go/configuration/src/configuration"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var logLevel = configuration.GetPropertyValue("LOG_LEVEL")

func SetLogOutput(fileName string) {
	log.SetOutput(&lumberjack.Logger{
		Filename:   "/var/log/" + fileName,
		MaxSize:    500,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	})
}

func Debug(message ...interface{}) {
	if logLevel == "DEBUG" {
		log.Println(message...)
	}
}

func Info(message ...interface{}) {
	if logLevel == "INFO" || logLevel == "DEBUG" {
		log.Println(message...)
	}
}

func Error(message ...interface{}) {
	if logLevel == "ERROR" || logLevel == "INFO" || logLevel == "DEBUG" {
		log.Println(message...)
	}
}

func ErrorException(message interface{}, err error) {
	if logLevel == "ERROR" {
		log.Println(message, err)
	}
}

func Fatal(v ...interface{}) {
	Error(fmt.Sprint(v...))
}
