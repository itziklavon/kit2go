package redis_helper

import (
	"encoding/json"

	"github.com/itziklavon/kit2go/configuration/src/configuration"
	"github.com/itziklavon/kit2go/general-log/src/general_log"
	"github.com/itziklavon/kit2go/http-client-helper/src/http_client_helper"
)

var DISCOVERY_URL = configuration.GetPropertyValue("DISCOVERY_URL")

var redisConnections map[int]RedisSessionHelper

type RedisData struct {
	BrandId int    `json:"brandId"`
	Uri     string `json:"uri"`
}

func GetRedisConnection(brandId int) RedisSessionHelper {
	if len(redisConnections) == 0 || redisConnections[brandId] == (RedisSessionHelper{}) {
		initMap()
	}
	value, ok := redisConnections[brandId]
	if ok {
		general_log.Debug("Key found value is: ", value)
	} else {
		general_log.Debug("Key not found")
		initMap()
	}
	return redisConnections[brandId]
}

func GetBrandSet() []int {
	if len(redisConnections) == 0 {
		initMap()
	}
	var brandSet []int
	for i := range redisConnections {
		brandSet = append(brandSet, i)
	}
	return brandSet
}

func initMap() {
	redisConnections = make(map[int]RedisSessionHelper)
	url := DISCOVERY_URL + "discovery-web/brand/services/REDIS"
	httpResponse := http_client_helper.GETBody(url, nil)
	var arr []RedisData
	_ = json.Unmarshal([]byte(httpResponse), &arr)
	for i := 0; i < len(arr); i = i + 1 {
		redisConnections[arr[i].BrandId] = RedisSessionHelper{host: arr[i].Uri}
	}
}
