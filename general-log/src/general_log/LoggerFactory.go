package general_log

import (
	"fmt"
	"configuration"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
	"log"
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

func Debug(message string) {
	if logLevel == "DEBUG" {
		log.Println(message)
	}
}

func Info(message string) {
	if logLevel == "INFO" || logLevel == "DEBUG" {
		log.Println(message)
	}
}

func Error(message string) {
	if logLevel == "ERROR" || logLevel == "INFO" || logLevel == "DEBUG" {
		log.Println(message)
	}
}

func ErrorException(message string, err error) {
	if logLevel == "ERROR" {
		log.Println(message, err)
	}
}

func Fatal(v ...interface{}) {
	Error(fmt.Sprint(v...))
}
