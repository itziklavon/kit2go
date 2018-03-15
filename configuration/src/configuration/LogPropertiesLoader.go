package configuration

import (
	"runtime"
	"strings"

	"log"
)

var logConf = "/opt/conf/log.properties"
var defaultLogConf = "log.properties"

type LogProperties struct {
	ready                   bool
	systemLogConfiguration  AppConfigProperties
	defaultLogConfiguration AppConfigProperties
}

var myLogConfiguration = LogProperties{ready: false}

func GetLogPropertyValue(key string) string {
	if !myLogConfiguration.ready {
		props, err := ReadPropertiesFile(logConf)
		if err != nil {
			log.Println(":GetPropertyValue: Error while reading properties file", err)
			return ""
		}
		_, filename, _, ok := runtime.Caller(0)
		if !ok {
			log.Println(":getDefaultValue: No caller information")
		}
		filename = strings.Replace(filename, "LogPropertiesLoader.go", "", 1) + defaultLogConf
		defaultProps := getDefaultValue(filename)
		myLogConfiguration = LogProperties{systemLogConfiguration: props, defaultLogConfiguration: defaultProps, ready: true}
	}
	if len(myLogConfiguration.systemLogConfiguration[key]) == 0 {
		return myLogConfiguration.defaultLogConfiguration[key]
	}
	return myLogConfiguration.systemLogConfiguration[key]

}

func getLogDefaultValue(key string) AppConfigProperties {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Println(":getDefaultValue: No caller information")
	}
	filename = strings.Replace(filename, "LogPropertiesLoader.go", "", 1) + defaultLogConf
	log.Println(filename)
	props, err := ReadPropertiesFile(filename)
	if err != nil {
		log.Println(":getDefaultValue: Error while reading properties file", err)
		return nil
	}
	return props
}
