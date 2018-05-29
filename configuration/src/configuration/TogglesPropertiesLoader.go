package configuration

import (
	"fmt"
	"runtime"
	"strings"

	"log"
)

var togglesConf = "/opt/conf/toggles.properties"
var defaultTogglesConf = "toggles.properties"

type TogglesProperties struct {
	ready                       bool
	systemTogglesConfiguration  AppConfigProperties
	defaultTogglesConfiguration AppConfigProperties
}

var myTogglesConfiguration = TogglesProperties{ready: false}

func GetTogglesPropertyValue(key string) string {
	if !myTogglesConfiguration.ready {
		props, err := ReadPropertiesFile(togglesConf)
		if err != nil {
			log.Println(":GetTogglesPropertyValue: Error while reading properties file", err)
			return ""
		}
		_, filename, _, ok := runtime.Caller(0)
		if !ok {
			log.Println(":GetTogglesPropertyValue: No caller information")
		}
		filename = strings.Replace(filename, "TogglesPropertiesLoader.go", "", 1) + defaultTogglesConf
		defaultProps := getDefaultTogglesValue(filename)
		myTogglesConfiguration = TogglesProperties{systemTogglesConfiguration: props, defaultTogglesConfiguration: defaultProps, ready: true}
		log.Println(":getDefaultValue: extracted conf are", myTogglesConfiguration)

	}
	if len(myTogglesConfiguration.systemTogglesConfiguration[key]) == 0 {
		return myTogglesConfiguration.defaultTogglesConfiguration[key]
	}
	return myTogglesConfiguration.systemTogglesConfiguration[key]

}

func getDefaultTogglesValue(key string) AppConfigProperties {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Println(":getDefaultValue: No caller information")
	}
	filename = strings.Replace(filename, "TogglesPropertiesLoader.go", "", 1) + defaultConf
	fmt.Println(filename)
	props, err := ReadPropertiesFile(filename)
	if err != nil {
		log.Println(":getDefaultValue: Error while reading properties file", err)
		return nil
	}
	return props
}
