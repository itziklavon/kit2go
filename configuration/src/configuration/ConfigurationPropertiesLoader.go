package configuration

import (
	"fmt"
	"runtime"
	"strings"

	"log"
)

var conf = "/opt/conf/configuration.properties"
var defaultConf = "configuration.properties"

type ConfiguraionProperties struct {
	ready                bool
	systemConfiguration  AppConfigProperties
	defaultConfiguration AppConfigProperties
}

var myConfiguration = ConfiguraionProperties{ready: false}

func GetPropertyValue(key string) string {
	if !myConfiguration.ready {
		props, err := ReadPropertiesFile(conf)
		if err != nil {
			log.Println(":GetPropertyValue: Error while reading properties file", err)
			return ""
		}
		_, filename, _, ok := runtime.Caller(0)
		if !ok {
			log.Println(":getDefaultValue: No caller information")
		}
		filename = strings.Replace(filename, "ConfigurationPropertiesLoader.go", "", 1) + defaultConf
		defaultProps := getDefaultValue(filename)
		myConfiguration = ConfiguraionProperties{systemConfiguration: props, defaultConfiguration: defaultProps, ready: true}
	}
	if len(myConfiguration.systemConfiguration[key]) == 0 {
		return myConfiguration.defaultConfiguration[key]
	}
	return myConfiguration.systemConfiguration[key]

}

func getDefaultValue(key string) AppConfigProperties {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Println(":getDefaultValue: No caller information")
	}
	filename = strings.Replace(filename, "ConfigurationPropertiesLoader.go", "", 1) + defaultConf
	fmt.Println(filename)
	props, err := ReadPropertiesFile(filename)
	if err != nil {
		log.Println(":getDefaultValue: Error while reading properties file", err)
		return nil
	}
	return props
}
