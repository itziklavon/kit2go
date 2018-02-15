package configuration

import (
	"runtime"
	"strings"

	"log"
)

var mongoConf = "/opt/conf/mongo.properties"
var defaultMongoConf = "mongo.properties"

type MongoConfiguraionProperties struct {
	ready                     bool
	mongoSystemConfiguration  AppConfigProperties
	mongoDefaultConfiguration AppConfigProperties
}

var myMongoConfiguration = MongoConfiguraionProperties{ready: false}

func GetMongoPropertyValue(key string) string {
	if !myConfiguration.ready {
		props, err := ReadPropertiesFile(conf)
		if err != nil {
			log.Println(":GetMongoPropertyValue: Error while reading properties file", err)
			return ""
		}
		_, filename, _, ok := runtime.Caller(0)
		if !ok {
			log.Println(":GetMongoPropertyValue: No caller information")
		}
		filename = strings.Replace(filename, "ConfigurationPropertiesLoader.go", "", 1) + defaultConf
		defaultProps := getMongoDefaultValue(filename)
		myMongoConfiguration = MongoConfiguraionProperties{mongoSystemConfiguration: props, mongoDefaultConfiguration: defaultProps, ready: true}
	}
	if len(myMongoConfiguration.mongoSystemConfiguration[key]) == 0 {
		return myMongoConfiguration.mongoDefaultConfiguration[key]
	}
	return myMongoConfiguration.mongoSystemConfiguration[key]
}

func getMongoDefaultValue(key string) AppConfigProperties {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Println(":getMongoDefaultValue: No caller information")
	}
	filename = strings.Replace(filename, "ConfigurationPropertiesLoader.go", "", 1) + defaultConf
	log.Println(":getMongoDefaultValue: file name is:" + filename)
	props, err := ReadPropertiesFile(filename)
	if err != nil {
		log.Println(":getMongoDefaultValue: Error while reading properties file", err)
		return nil
	}
	return props
}
